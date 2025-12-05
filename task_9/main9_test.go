package main

import (
	"math"
	"testing"
)

func TestProcessNumbers(t *testing.T) {
	in := make(chan uint8, 3)
	out := make(chan float64, 3)

	go ProcessNumbers(in, out)

	in <- 1
	in <- 2
	in <- 3
	close(in)

	got := []float64{}
	for x := range out {
		got = append(got, x)
	}

	want := []float64{1.0, 8.0, 27.0}
	if len(got) != len(want) {
		t.Fatalf("ожидалось %d элементов, получено %d", len(want), len(got))
	}
	for i, v := range got {
		if math.Abs(v-want[i]) > 1e-9 {
			t.Errorf("элемент %d: получено %f, ожидалось %f", i, v, want[i])
		}
	}
}
