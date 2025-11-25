package main

import (
	"fmt"
	"reflect"
)

func main() {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//1
	fmt.Println("\n Шаг 1: Начальные значения: ", originalSlice)
	originalSlice = setRandomValues(originalSlice)

	fmt.Println("Случайные значения: ", originalSlice)

	//2
	newSlice := sliceExample(originalSlice)
	fmt.Println("\n Шаг 2: Чётные числа: ", newSlice)

	//3
	value := 2
	fmt.Printf("\n Шаг 3: Число %v добавлено %v", value, addElements(newSlice, value))

	//4
	copyS := copySlice(originalSlice)
	originalSlice = append(originalSlice, value)
	equal := reflect.DeepEqual(originalSlice, copyS) // → bool

	fmt.Printf("\n Шаг 4: Копия %v, равны: %v", copyS, equal)

	//5
	index := 2
	removeSl, err := removeElement(copyS, index)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n Шаг 5: Из слайса %v, удаляем элемент по индесу: %v, элемент %v, получаем %v", copyS, index, copyS[index], removeSl)
}
