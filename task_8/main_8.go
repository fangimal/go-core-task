package main

import (
	"fmt"
	"time"
)

func main() {
	var wg CustomWaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Горутина 1 готова")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Горутина 2 готова")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(150 * time.Millisecond)
		fmt.Println("Горутина 3 готова")
	}()

	fmt.Println("Ждём завершения...")
	wg.Wait()
	fmt.Println("Все горутины завершены!")
}
