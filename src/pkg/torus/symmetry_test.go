package torus

import (
	"testing"
	"reflect"
)

type BlockDimList struct {
	t   Torus
	p   []Point
	dim Torus
}

var L = []BlockDimList{
	{Torus{10, 10}, []Point{{0, 0}}, Torus{10, 10}},
	{Torus{10, 10}, []Point{{5, 5}}, Torus{10, 10}},
	{Torus{10, 10}, []Point{{0, 5}}, Torus{10, 5}},
	{Torus{10, 10}, []Point{{5, 0}}, Torus{5, 10}},
	{Torus{10, 10}, []Point{{5, 5}, {0, 5}}, Torus{10, 5}},
	{Torus{10, 10}, []Point{{5, 5}, {5, 0}}, Torus{5, 10}},
	{Torus{10, 10}, []Point{{5, 5}, {5, 0}, {0, 5}}, Torus{5, 5}},
}

func TestBlockDim(t *testing.T) {
	for _, l := range L {
		e := l.t.BlockDim(l.p)
		if !e.Equal(l.dim) {
			t.Errorf("BlockDim %v expected %v for %v %v", e, l.dim, l.t, l.p)
		}
	}
}

type ReduceList struct {
	t    Torus
	good bool
	n    int
	e    Point
	p    []Point
}

var LR = []ReduceList{
	{Torus{7, 7}, true, 7, Point{1, -2}, []Point{{0, 0}, {1, 5}}},
	{Torus{7, 7}, true, 7, Point{2, 1}, []Point{{2, 1}, {4, 2}}},
	{Torus{7, 7}, true, 7, Point{2, -1}, []Point{{0, 0}, {2, -1}}},
	{Torus{7, 7}, true, 7, Point{2, -1}, []Point{{0, 0}, {-2, 1}}},
	{Torus{7, 7}, true, 7, Point{2, -1}, []Point{{0, 0}, {-2, 1}}},
	{Torus{7, 7}, false, 1, Point{0, 0}, []Point{{0, 0}, {0, 0}}},
	{Torus{196, 196}, true, 4, Point{98, -49}, []Point{{22, 58}, {120, 9}}},
	{Torus{96, 80}, true, 4, Point{24, -20}, []Point{{24, -20}, {48, 40}}},
}

func xTestReduce(t *testing.T) {
	for _, l := range LR {
		p1 := l.p[0]
		for _, p := range l.p[1:] {
			got, good := l.t.ShiftReduce(p1, p, 32)
			if good != l.good || !l.t.PointEqual(got, l.e) {
				t.Error("ShiftReduce mismatch", l.p[0], p, got, good, l)
			}
			nt := l.t.TranslationLen(got)
			if nt != l.n {
				t.Error("TranslationLen mismatch", l.t, p, nt, l.n, l)
			}
		}
	}
}

type RRList struct {
	t   Torus
	dim Torus
	e   []Point
	re  []Point
	red []Point
}

var RR = []RRList{
	//	{Torus{64, 64}, Torus{16, 16}, []Point{{32, 32}}, []Point{}, []Point{{16, 0}, {0, 16}, {32, 32}}},
	//	{Torus{102, 129}, Torus{102, 129}, []Point{{34, 43}}, []Point{{34, 43}}, []Point{{34, 43}, {34, 43}, {34, 43}}},
	{Torus{96, 80}, Torus{96, 80}, []Point{{24, -20}}, []Point{{24, -20}},
		[]Point{{24, -20}, {24, -20}, {48, 40}, {48, 40}, {24, -20}, {24, -20}, {24, -20}, {24, -20}, {48, 40}, {48, 40}, {24, -20}, {24, -20}, {24, -20}, {24, -20}, {48, 40}, {48, 40}, {24, -20}, {24, -20}, {48, 40}, {48, 40}, {24, -20}, {24, -20}, {24, -20}, {24, -20}}},
}

func TestReduceReduce(t *testing.T) {
	for _, l := range RR {
		dim := l.t
		e := dim.ReduceReduce(l.red)
		if !reflect.DeepEqual(e, l.e) {
			t.Error("Mismatched ", e, l.e, " data ", l)
		}
		dim = dim.BlockDim(l.red)
		for i, p := range l.red {
			l.red[i] = dim.Donut(p)
		}
		e = dim.ReduceReduce(l.red)
		if !reflect.DeepEqual(e, l.re) {
			t.Error("Reduction mismatched ", e, l.re, " data ", l)
		}
	}
}
