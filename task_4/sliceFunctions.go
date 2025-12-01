package main

import "slices"

func uniqueElementsFromFirst(first, second []string) []string {
	newSlice := make([]string, 0, len(first))

	for _, value := range first {
		if !slices.Contains(second, value) {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}
