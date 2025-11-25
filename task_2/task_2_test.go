package main

import (
	"reflect"
	"testing"
)

func TestSetRandomValues(t *testing.T) {
	original := []int{-1, -1, -1}

	result := setRandomValues(original)

	if len(result) != 3 {
		t.Errorf("длина изменилась: %d", len(result))
	}

	for i, v := range result {
		if v < 1 || v > 100 {
			t.Errorf("элемент %d = %d вне диапазона [1,100]", i, v)
		}
	}

	if result[0] == -1 || result[1] == -1 || result[2] == -1 {
		t.Error("значения не обновились")
	}
}

// TestSliceExample фильтрация чётных чисел.
func TestSliceExample(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"пустой", []int{}, []int{}},
		{"без чётных", []int{1, 3, 5}, []int{}},
		{"все чётные", []int{2, 4, 6}, []int{2, 4, 6}},
		{"смешанный", []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{"с отрицательными", []int{-2, -1, 0, 1, 2}, []int{-2, 0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceExample(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("sliceExample(%v) = %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}
}

// TestAddElements добавление элемента в конец.
func TestAddElements(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		value    int
		expected []int
	}{
		{"в пустой", []int{}, 5, []int{5}},
		{"в непустой", []int{1, 2}, 3, []int{1, 2, 3}},
		{"нулевое значение", []int{0}, 0, []int{0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addElements(tt.input, tt.value)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("addElements(%v, %d) = %v, expected %v", tt.input, tt.value, got, tt.expected)
			}
		})
	}
}

// TestCopySlice копия независима от оригинала.
func TestCopySlice(t *testing.T) {
	orig := []int{1, 2, 3, 4, 5}
	copied := copySlice(orig)

	// Меняем оригинал
	orig[0] = 999
	orig = append(orig, 100)

	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(copied, expected) {
		t.Errorf("копия изменилась: %v, ожидалось %v", copied, expected)
	}

	// Проверим, что это разные слайсы (по ёмкости или адресу — опционально)
	if &copied[0] == &orig[0] {
		t.Error("копия и оригинал используют одну память!")
	}
}

// TestRemoveElement удаление по индексу.
func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		index       int
		expected    []int
		expectError bool
	}{
		{"удалить из середины", []int{1, 2, 3, 4}, 1, []int{1, 3, 4}, false},
		{"удалить первый", []int{1, 2, 3}, 0, []int{2, 3}, false},
		{"удалить последний", []int{1, 2, 3}, 2, []int{1, 2}, false},
		{"один элемент", []int{42}, 0, []int{}, false},
		{"ошибка: индекс < 0", []int{1, 2}, -1, nil, true},
		{"ошибка: индекс >= len", []int{1, 2}, 2, nil, true},
		{"ошибка: пустой слайс", []int{}, 0, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := removeElement(tt.input, tt.index)

			if tt.expectError {
				if err == nil {
					t.Error("ожидалась ошибка, но её нет")
				}
				return
			}

			if err != nil {
				t.Fatalf("неожиданная ошибка: %v", err)
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("removeElement(%v, %d) = %v, expected %v", tt.input, tt.index, got, tt.expected)
			}
		})
	}
}
