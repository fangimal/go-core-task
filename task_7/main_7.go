package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Создаём 3 канала
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Запускаем горутины-источники
	go func() {
		defer close(ch1)
		for _, x := range []int{1, 3, 5} {
			ch1 <- x
		}
	}()

	go func() {
		defer close(ch2)
		for _, x := range []int{2, 4} {
			ch2 <- x
		}
	}()

	go func() {
		defer close(ch3)
		for _, x := range []int{10, 20} {
			ch3 <- x
		}
	}()

	// Сливаем
	merged := Merge(ch1, ch2, ch3)

	// Собираем результат
	result := []int{}
	for x := range merged {
		result = append(result, x)
	}

	fmt.Println("Слияние:", result)
	// Пример вывода: [1 2 3 10 4 5 20] (порядок может отличаться)
}

// Merge сливает N каналов в один.
// Входные каналы должны быть закрыты отправителями (иначе merged никогда не закроется).
// Результат — канал, который закроется, когда все входные каналы будут закрыты.
func Merge(channels ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		// 1. Преобразуем слайс каналов в []reflect.SelectCase
		// Каждый case — операция чтения (SelectRecv) из канала
		var cases []reflect.SelectCase
		for _, ch := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,  // хотим читать
				Chan: reflect.ValueOf(ch), // канал для чтения
			})
		}

		// 2. Пока есть каналы — слушаем их

		for len(cases) > 0 {
			// 3. reflect.Select блокируется, пока один из каналов не будет готов
			// Возвращает:
			//   chosen — индекс канала в cases,
			//   value — переданное значение (reflect.Value),
			//   ok     — true, если канал не закрыт
			chosen, value, ok := reflect.Select(cases)

			if ok {
				// 4. Канал жив — отправляем значение в выходной канал
				out <- int(value.Int()) // Int() — потому что канал int
			} else {
				// 5. Канал закрыт — удаляем его из cases
				// Удаляем элемент по индексу chosen (без аллокаций)
				cases[chosen] = cases[len(cases)-1] // копируем последний на место chosen
				cases = cases[:len(cases)-1]        // уменьшаем длину
			}
		}
	}()

	return out
}
