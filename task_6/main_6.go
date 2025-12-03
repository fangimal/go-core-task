package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count := 5
	ch := RandomGenerator(count, time.Now().UnixNano())
	nums := make([]int, 0, count)

	for x := range ch {
		nums = append(nums, x)
	}

	fmt.Printf("Datas: %v\n", nums)
}

func RandomGenerator(n int, seed int64) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		r := rand.New(rand.NewSource(seed))
		for i := 0; i < n; i++ {
			value := r.Intn(100)
			ch <- value
		}
	}()

	return ch
}
