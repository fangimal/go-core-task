package main

import (
	"slices"
	"testing"
)

func TestIntEntrySlice(t *testing.T) {
	tests := []struct {
		name      string
		first     []int
		second    []int
		wantBool  bool
		wantSlice []int
	}{
		{
			name:      "обычный случай",
			first:     []int{65, 3, 58, 678, 64},
			second:    []int{64, 2, 3, 43},
			wantBool:  true,
			wantSlice: []int{3, 64},
		},
		{
			name:      "пустой second",
			first:     []int{65, 3},
			second:    []int{},
			wantBool:  false,
			wantSlice: []int{},
		},
		{
			name:      "пустой first",
			first:     []int{},
			second:    []int{1},
			wantBool:  false,
			wantSlice: []int{},
		},
		{
			name:      "полное пересечение",
			first:     []int{65, 3, 58, 678, 64},
			second:    []int{65, 3, 58, 678, 64},
			wantBool:  true,
			wantSlice: []int{65, 3, 58, 678, 64},
		},
		{
			name:      "дубликаты в first",
			first:     []int{65, 3, 58, 678, 65},
			second:    []int{65, 2, 3, 43},
			wantBool:  true,
			wantSlice: []int{65, 3, 65},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, got := intEntrySlice(tt.first, tt.second)
			if !slices.Equal(got, tt.wantSlice) || ok != tt.wantBool {
				t.Errorf("intEntrySlice(%v, %v) = (%v, %v), want (%v, %v)",
					tt.first, tt.second, ok, got, tt.wantBool, tt.wantSlice)
			}
		})
	}
}
