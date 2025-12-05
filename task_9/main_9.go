package main

import (
	"fmt"
)

func main() {
	input := make(chan uint8)
	output := make(chan float64)

	go ProcessNumbers(input, output)

	go func() {
		defer close(input)
		for _, x := range []uint8{2, 3, 4} {
			input <- x
		}
	}()

	fmt.Println("Результаты:")
	for cube := range output {
		fmt.Printf("%.1f\n", cube)
	}
}

func ProcessNumbers(in <-chan uint8, out chan<- float64) {
	defer close(out)

	for x := range in {
		fx := float64(x)
		cube := fx * fx * fx
		out <- cube
	}
}
