package main

import (
	"reflect"
	"testing"
)

func TestRandomGenerator(t *testing.T) {
	// При одинаковом seed — последовательность должна быть одинаковой
	ch1 := RandomGenerator(3, 123)
	ch2 := RandomGenerator(3, 123)

	var result1, result2 []int
	for x := range ch1 {
		result1 = append(result1, x)
	}
	for x := range ch2 {
		result2 = append(result2, x)
	}

	if !reflect.DeepEqual(result1, result2) {
		t.Errorf("ожидалось одинаковые последовательности, получено %v и %v", result1, result2)
	}

	if len(result1) != 3 {
		t.Errorf("ожидалось 3 числа, получено %d", len(result1))
	}

	//t.Logf("Получено: %v, %v ", result1, result2)
}
