package engine

import (
	"log"
	"testing"
	"bugnuts/replay"
	"fmt"
	"os"
)

func TestEngine(t *testing.T) {

	match, err := replay.Load("testdata/0.replay")

	if err != nil || match == nil {
		log.Panicf("Error loading replay: %v", err)
	}
	m := match.GetMap()

	g := NewGame(&match.GameInfo, m)

	tout := g.Replay(match.Replay, 0, match.GameLength, true)

	for i := range tout[0] {
		filein := fmt.Sprint("testdata/0.bot", i, ".input")
		fileout := filein + ".tmp"
		out, err := os.Create(fileout)
		defer out.Close()
		if err != nil {
			log.Panic("open failed for ", fileout, ":", err)
		}
		fmt.Fprintf(out, "turn 0\n%v\nready\n", g.GameInfo)
		for turn := range tout {
			if len(tout[turn][i].A) > 0 && turn < match.GameLength {
				fmt.Fprint(out, "turn ", turn+1, "\n")
			} else {
				fmt.Fprint(out, "end\n")
			}
			fmt.Fprint(out, tout[turn][i], "\ngo\n")
			if len(tout[turn][i].A) == 0 {
				break
			}
		}
	}
}

// BenchmarkEngine times turn generation for a replay file.
func BenchmarkEngine(b *testing.B) {

	match, err := replay.Load("testdata/0.replay")
	if err != nil || match == nil {
		log.Panicf("Error loading replay: %v", err)
	}
	m := match.GetMap()

	for i := 0; i < b.N; i++ {
		g := NewGame(&match.GameInfo, m)
		g.Replay(match.Replay, 0, match.GameLength, false)
	}
}

// BenchmarkEngineOrdered times the generation in cannonical order to reproduce the 
// python input files.
func BenchmarkEngineOrdered(b *testing.B) {

	match, err := replay.Load("testdata/0.replay")
	if err != nil || match == nil {
		log.Panicf("Error loading replay: %v", err)
	}
	m := match.GetMap()

	for i := 0; i < b.N; i++ {
		g := NewGame(&match.GameInfo, m)
		g.Replay(match.Replay, 0, match.GameLength, true)
	}
}
