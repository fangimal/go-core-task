// hasher_test.go
package main

import (
	"reflect"
	"testing"
)

// TestInspectVariables проверяет, что функция корректно определяет типы и сохраняет значения.
func TestInspectVariables(t *testing.T) {
	// Подготовка
	num := 42
	pi := 3.14
	name := "Golang"
	ok := true
	c := complex64(1 + 2i)

	// Действие
	infos := inspectVariables(num, pi, name, ok, c)

	// Проверка
	if len(infos) != 5 {
		t.Fatalf("ожидалось 5 элементов, получено %d", len(infos))
	}

	cases := []struct {
		index     int
		wantType  string
		wantValue interface{}
	}{
		{0, "int", num},
		{1, "float64", pi},
		{2, "string", name},
		{3, "bool", ok},
		{4, "complex64", c},
	}

	for _, tc := range cases {
		got := infos[tc.index]
		if got.Type != tc.wantType {
			t.Errorf("элемент %d: ожидался тип %q, получен %q", tc.index, tc.wantType, got.Type)
		}
		if !reflect.DeepEqual(got.Value, tc.wantValue) {
			t.Errorf("элемент %d: ожидалось значение %v, получено %v", tc.index, tc.wantValue, got.Value)
		}
	}
}

// TestValuesToStrings проверяет преобразование значений в строки.
func TestValuesToStrings(t *testing.T) {
	infos := []TypeValue{
		{Value: 42, Type: "int"},
		{Value: 3.14, Type: "float64"},
		{Value: "Golang", Type: "string"},
		{Value: true, Type: "bool"},
		{Value: complex64(1 + 2i), Type: "complex64"},
	}

	got := valuesToStrings(infos)
	want := []string{"42", "3.14", "Golang", "true", "(1+2i)"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ожидалось %v, получено %v", want, got)
	}
}

// TestHashRunesWithSaltInMiddle проверяет детерминированность и длину хэша.
func TestHashRunesWithSaltInMiddle(t *testing.T) {
	runes := []rune("test")
	salt := "go-2024"

	hash1 := hashRunesWithSaltInMiddle(runes, salt)
	hash2 := hashRunesWithSaltInMiddle(runes, salt)

	// Хэш должен быть детерминированным
	if hash1 != hash2 {
		t.Errorf("хэш не детерминированный: %s != %s", hash1, hash2)
	}

	// Длина SHA256 в hex — всегда 64 символа
	if len(hash1) != 64 {
		t.Errorf("длина хэша %d, ожидалось 64", len(hash1))
	}

	// Проверим, что разные входы → разные хэши
	hash3 := hashRunesWithSaltInMiddle([]rune("test2"), salt)
	if hash1 == hash3 {
		t.Errorf("разные входы дали одинаковый хэш: %s", hash1)
	}
}

// ExampleHashRunesWithSaltInMiddle — пример для go doc и godoc.org
func ExamplehashRunesWithSaltInMiddle() {
	runes := []rune("42 Golang")
	salt := "go-2024"
	hash := hashRunesWithSaltInMiddle(runes, salt)
	// Обрезаем хэш для примера
	println(hash[:8] + "...")
}
