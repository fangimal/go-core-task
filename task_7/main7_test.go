package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	// Тест 1: 2 канала
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 1)
	ch1 <- 1
	ch1 <- 3
	close(ch1)
	ch2 <- 2
	close(ch2)

	merged := Merge(ch1, ch2)
	got := []int{}
	for x := range merged {
		got = append(got, x)
	}

	// Порядок не гарантируется, но множество должно быть {1,2,3}
	wantSet := map[int]bool{1: true, 2: true, 3: true}
	if len(got) != 3 {
		t.Fatalf("ожидалось 3 элемента, получено %d", len(got))
	}
	for _, x := range got {
		if !wantSet[x] {
			t.Errorf("лишнее значение %d", x)
		}
		delete(wantSet, x)
	}
	if len(wantSet) != 0 {
		t.Errorf("не хватает значений: %v", wantSet)
	}

	// Тест 2: пустые каналы
	empty1 := make(chan int)
	close(empty1)
	empty2 := make(chan int)
	close(empty2)

	merged2 := Merge(empty1, empty2)
	got2 := []int{}
	for x := range merged2 {
		got2 = append(got2, x)
	}
	if len(got2) != 0 {
		t.Errorf("ожидался пустой слайс, получено %v", got2)
	}

	// Тест 3: один канал
	single := make(chan int, 1)
	single <- 42
	close(single)

	merged3 := Merge(single)
	got3 := []int{}
	for x := range merged3 {
		got3 = append(got3, x)
	}
	if !reflect.DeepEqual(got3, []int{42}) {
		t.Errorf("ожидалось [42], получено %v", got3)
	}
}
