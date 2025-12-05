package main

import (
	"slices"
	"testing"
)

func TestSliceFunctions(t *testing.T) {
	tests := []struct {
		name   string
		first  []string
		second []string
		want   []string
	}{
		{
			name:   "обычный случай",
			first:  []string{"apple", "banana", "cherry"},
			second: []string{"banana", "date"},
			want:   []string{"apple", "cherry"},
		},
		{
			name:   "пустой second",
			first:  []string{"a", "b"},
			second: []string{},
			want:   []string{"a", "b"},
		},
		{
			name:   "пустой first",
			first:  []string{},
			second: []string{"x"},
			want:   []string{},
		},
		{
			name:   "полное пересечение",
			first:  []string{"a", "b"},
			second: []string{"b", "a"},
			want:   []string{},
		},
		{
			name:   "дубликаты в first",
			first:  []string{"a", "a", "b"},
			second: []string{"a"},
			want:   []string{"b"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniqueElementsFromFirst(tt.first, tt.second)
			if !slices.Equal(got, tt.want) {
				t.Errorf("uniqueElementsFromFirst(%v, %v) = %v, expected %v",
					tt.first, tt.second, got, tt.want)
			}
		})
	}
}
