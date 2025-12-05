package main

import (
	"fmt"
)

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	slice3 := uniqueElementsFromFirst(slice1, slice2)

	fmt.Printf("Результат: %v", slice3)
}
