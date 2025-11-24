package main

import (
	"fmt"
	"strings"
)

func main() {
	//1.
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	//2.
	infos := inspectVariables(
		numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)

	fmt.Println("\n Шаг 2: Типы переменных")
	for i, v := range infos {
		fmt.Printf(" [%d] тип: %s, значение: %v\n", i, v.Type, v.Value)
	}

	//3.
	fmt.Println("\n Шаг 3: Преобразование значений в строки")
	strs := valuesToStrings(infos)
	fmt.Printf("Строки: %v\n", strs)

	joined := strings.Join(strs, " ")
	fmt.Printf("Объединённая строка: %q\n", joined)

	//4.
	fmt.Println("\n Шаг 4: Преобразование в срез рун")
	runes := stringToRunes(joined)
	fmt.Printf("Длина среза рун: %d\n", len(runes))
	fmt.Printf("Runes:   %q\n", string(runes))

	// === ШАГ 5: Хэширование с солью в середине ===
	fmt.Println("\n Шаг 5: Хэширование с солью 'go-2024' в середине")

	salt := "go-2024"
	hash := hashRunesWithSaltInMiddle(runes, salt)

	// Выводим результат
	fmt.Printf("Хэш SHA256: %s\n", hash)
	fmt.Printf("Длина хэша: %d символов (ожидаемо: 64)\n", len(hash))

	// 🔍 Для проверки: покажем, как выглядит строка с солью
	s := string(runes)
	mid := len(s) / 2
	saltedPreview := s[:mid] + "[" + salt + "]" + s[mid:]
	fmt.Printf("Строка с солью (в скобках для наглядности):\n%q\n", saltedPreview)
}
