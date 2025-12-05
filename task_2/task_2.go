package main

import (
	"errors"
	"math/rand"
)

func setRandomValues(originalSlice []int) []int {
	for i, value := range originalSlice {
		value = rand.Intn(100) + 1
		originalSlice[i] = value
	}
	return originalSlice
}

func sliceExample(numbers []int) []int {
	newSlice := make([]int, 0, len(numbers))
	for _, value := range numbers {
		if value%2 == 0 {
			newSlice = append(newSlice, value)
		}
	}

	return newSlice
}

func addElements(numbers []int, value int) []int {
	return append(numbers, value)
}

func copySlice(numbers []int) []int {
	return append([]int(nil), numbers...)
}

func removeElement(numbers []int, index int) ([]int, error) {
	if index < 0 || index >= len(numbers) {
		return nil, errors.New("index out of range")
	}
	result := make([]int, len(numbers)-1)
	copy(result, numbers[:index])
	copy(result[index:], numbers[index+1:])
	return result, nil
}
