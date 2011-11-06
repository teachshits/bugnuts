package main

import (
	"testing"
	"os"
	"log"
	"time"
	"fmt"
	"strconv"
)

var maps = []string{
	"maze_02p_01",
	"maze_02p_02",
	"maze_03p_01",
	"maze_04p_01",
	"maze_04p_02",
	"maze_05p_01",
	"maze_06p_01",
	"maze_07p_01",
	"maze_08p_01",
	"mmaze_02p_01",
	"mmaze_02p_02",
	"mmaze_03p_01",
	"mmaze_04p_01",
	"mmaze_04p_02",
	"mmaze_05p_01",
	"mmaze_07p_01",
	"mmaze_08p_01",
	"random_walk_02p_01",
	"random_walk_02p_02",
	"random_walk_03p_01",
	"random_walk_03p_02",
	"random_walk_04p_01",
	"random_walk_04p_02",
	"random_walk_05p_01",
	"random_walk_05p_02",
	"random_walk_06p_01",
	"random_walk_06p_02",
	"random_walk_07p_01",
	"random_walk_07p_02",
	"random_walk_08p_01",
	"random_walk_08p_02",
	"random_walk_09p_01",
	"random_walk_09p_02",
	"random_walk_10p_01",
	"random_walk_10p_02",
}


func TestMapFill(t *testing.T) {
	file := "testdata/fill2.map" // fill.2 Point{r:4, c:5}
	m, err := MapLoadFile(file)

	if err != os.EOF {
		t.Errorf("Read failed for %s: %v", file, err)
	} else if m == nil {
		t.Errorf("Invalid load of map m == nil")
	}

	// log.Printf("%v", m) // TODO test String() func round trip.
	l := make(map[Location]int, 0)
	for _, hill := range m.HillLocations(-1) {
		l[hill] = 1
	}

	fs, mQ, mD := MapFill(m, l, 1)
	log.Printf("SlowFill: mQ: %v mD: %v f::\n%v\n", mQ, mD, fs)
	ff, mQ, mD := MapFillSlow(m, l, 1)
	log.Printf("FastFill: mQ: %v mD: %v f::\n%v\n", mQ, mD, ff)
	
	

}

// Benchmark the version which does not maintain a seed array
// but allocates per fill
func BenchmarkMapFillAlloc(b *testing.B) {
	file := "testdata/maps/mmaze_05p_01.map"
	m, err := MapLoadFile(file)
	if m == nil || err != os.EOF {
		log.Panicf("Error reading %s: err %v map: %v", file, err, m)
	}


	l := make(map[Location]int, 40)

	for _, hill := range m.HillLocations(-1) {
		l[hill] = 1
	}

	for i := 0; i < b.N; i++ {
		MapFill(m, l, 1)
	}
}

// Benchmark not reusing the fill struct.
func BenchmarkMapFill(b *testing.B) {
	file := "testdata/maps/mmaze_05p_01.map"
	m, err := MapLoadFile(file)
	if m == nil || err != os.EOF {
		log.Panicf("Error reading %s: err %v map: %v", file, err, m)
	}

	l := make(map[Location]int, 40)

	for _, hill := range m.HillLocations(-1) {
		l[hill] = 1
	}
	f := m.NewFill()
	for i := 0; i < b.N; i++ {
		f.Reset()
		f.MapFill(m, l, 1)
	}
}

// Benchmark allocating fill + computing seed.
func BenchmarkMapFillSeed(b *testing.B) {
	file := "testdata/maps/mmaze_05p_01.map"
	m, err := MapLoadFile(file)
	if m == nil || err != os.EOF {
		log.Panicf("Error reading %s: err %v map: %v", file, err, m)
	}

	l := make(map[Location]int, 40)

	for _, hill := range m.HillLocations(-1) {
		l[hill] = 1
	}
	for i := 0; i < b.N; i++ {
		MapFillSeed(m, l, 1)
	}
}



func TestMapFillDist(t *testing.T) {
	out, _ := os.Create("tmp/dist.csv")
	defer out.Close()
	
	for _, name := range maps {
		filename := "testdata/maps/" + name + ".map"
		m, err := MapLoadFile(filename)
		if m == nil || err != os.EOF {
			log.Panicf("Error: failed to read %s: %v", filename, err)
		}
		for _, player := range []int{-1,0} {
			l := make(map[Location]int)
			for _, hill := range m.HillLocations(player) {
				l[hill] = 1
			}
			pre := time.Nanoseconds()
			f, mQ, mD := MapFillSlow(m, l, 1)
			post := time.Nanoseconds()
			ff, mQ, mD := MapFill(m, l, 1)
			postff := time.Nanoseconds()
			ffs, mQ, mD := MapFillSeed(m, l, 1)
			postffs := time.Nanoseconds()
			diff := 0
			for i, f := range f.Depth {
				if f != ff.Depth[i] || f != ffs.Depth[i] || ffs.Depth[i] != ff.Depth[i] {
					diff++	
				}
			}
			log.Printf("Fill: mQ:%3d mD: %3d %4.1f/%4.1f/%4.1f ms %d diffs player %d points %d %s", mQ, mD, float64(post-pre)/1000000, float64(postff-post)/1000000, float64(postffs-postff)/1000000, diff, player, len(l), name)

			// Generate histograms.
			empty := NewMap(m.Rows, m.Cols, 1)
			fe, _, mDe := MapFill(empty, l, 1)
			if mD > mDe {
				mDe = mD
			}

			histe := make([]int,mDe+1)
			hist := make([]int,mDe+1)
			for i, d := range f.Depth {
				hist[d]++
				histe[fe.Depth[i]]++
			}
			if player == 0 {
				for i, k := range hist {
					fmt.Fprintf(out, "\"%s\",%d,%d,%d\n", name, i, k, histe[i])
				}
			}
		}
	}
}


// Take a map and generate montecarlo ant densities...
func TestMonteCarloPathing(t *testing.T) {
	for _, name := range maps {
		filename := "testdata/maps/" + name + ".map"
		m, err := MapLoadFile(filename)
		if m == nil || err != os.EOF {
			log.Panicf("Error: failed to read %s: %v", filename, err)
		}

		lend := make(map[Location]int)
		for _, hill := range m.HillLocations(-1) {
			lend[hill] = 1
		}

		//lsrc := m.HillLocations(-1)
		lsrc := make([]Location,0,len(m.Grid))

		f,_,_ := MapFillSeed(m, lend, 1)

		for loc, item := range m.Grid {
			if item != WATER {
				lsrc = append(lsrc, Location(loc))
			}
		}
		
		d := 100000/len(lsrc)
		pre := time.Nanoseconds()
		paths := f.MontePathIn(m, lsrc, d, 1)
		post := time.Nanoseconds()
		log.Printf("Montecarlo %d paths in %.2f ms", d*len(lsrc), float64(post-pre)/1000000)

		str := ""
		// Write out in the annoying R image layout
		for c := 0; c < m.Cols; c++ {
			for r := m.Rows-1; r >= 0; r-- {
				loc := Location(r * m.Cols + c)
				if r != m.Rows - 1 {
					str += ","
				}
				str += strconv.Itoa(int(paths[loc]))
			}
			str += "\n"
		}
		out, _ := os.Create("tmp/" + name + ".csv")
		fmt.Fprint(out, str)
		out.Close()		
	}
}
			
func xTestP(t *testing.T) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j != i {
				for k := 0; k < 4; k++ {
					if k != i && k != j {
						for l := 0; l < 4; l++ {
							if l != i && l != j && l != k {
								log.Printf("%d,%d,%d,%d",i,j,k,l)
							}
						}
					}
				}
			}
		}
	}
}
					
	
