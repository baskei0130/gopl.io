package popcount

import (
	"testing"

	"gopl.io/chap2/2.6.2/popcount"
)

func TestPopCount1(t *testing.T) {
	var tests = []struct {
		x    uint64
		want int
	}{
		{uint64(0), 0},
		{uint64(3), 2},
		{uint64(4), 1},
		{uint64(63), 6},
	}

	for _, test := range tests {
		if got := PopCount1(test.x); got != test.want {
			t.Errorf("PopCount(%d) = %d; want %d", test.x, got, test.want)
		}
	}
}

func TestPopCount2(t *testing.T) {
	var tests = []struct {
		x    uint64
		want int
	}{
		{uint64(0), 0},
		{uint64(3), 2},
		{uint64(4), 1},
		{uint64(63), 6},
	}

	for _, test := range tests {
		if got := PopCount2(test.x); got != test.want {
			t.Errorf("PopCount(%d) = %d; want %d", test.x, got, test.want)
		}
	}
}

func TestPopCount3(t *testing.T) {
	var tests = []struct {
		x    uint64
		want int
	}{
		{uint64(0), 0},
		{uint64(3), 2},
		{uint64(4), 1},
		{uint64(63), 6},
	}

	for _, test := range tests {
		if got := PopCount3(test.x); got != test.want {
			t.Errorf("PopCount(%d) = %d; want %d", test.x, got, test.want)
		}
	}
}

func TestPopCount4(t *testing.T) {
	var tests = []struct {
		x    uint64
		want int
	}{
		{uint64(0), 0},
		{uint64(3), 2},
		{uint64(4), 1},
		{uint64(63), 6},
	}

	for _, test := range tests {
		if got := PopCount4(test.x); got != test.want {
			t.Errorf("PopCount(%d) = %d; want %d", test.x, got, test.want)
		}
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(uint64(i))
	}
}

func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount1(uint64(i))
	}
}
