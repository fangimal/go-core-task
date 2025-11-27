package main

import (
	"reflect"
	"testing"
)

func TestNewStringIntMap(t *testing.T) {
	m := NewStringIntMap()
	if m == nil {
		t.Fatal("NewStringIntMap вернул nil")
	}
	if m.data == nil {
		t.Fatal("внутренняя мапа не инициализирована")
	}
}

func TestAdd(t *testing.T) {
	m := NewStringIntMap()

	m.Add("a", 1)
	v, ok := m.Get("a")
	if !ok || v != 1 {
		t.Errorf("после Add(\"a\", 1): Get(\"a\") = (%d, %v), ожидалось (1, true)", v, ok)
	}

	// Обновление существующего ключа
	m.Add("a", 42)
	v, ok = m.Get("a")
	if !ok || v != 42 {
		t.Errorf("после обновления: Get(\"a\") = (%d, %v), ожидалось (42, true)", v, ok)
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)

	m.Remove("a")
	if m.Exists("a") {
		t.Error("ключ \"a\" всё ещё существует после Remove")
	}
	if !m.Exists("b") {
		t.Error("ключ \"b\" исчез после удаления \"a\"")
	}

	// Удаление несуществующего — не должно паниковать
	m.Remove("не_существует")
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)

	originalMap := m.Copy() // копия ДО изменений

	// Меняем оригинал
	m.Add("c", 3)
	m.Remove("a")

	// Копия должна остаться неизменной
	expected := map[string]int{"a": 1, "b": 2}
	if !reflect.DeepEqual(originalMap, expected) {
		t.Errorf("копия изменилась: %v, ожидалось %v", originalMap, expected)
	}

	// Убедимся, что это разные объекты в памяти
	newCopy := m.Copy()
	newCopy["d"] = 4
	if m.Exists("d") {
		t.Error("изменение копии повлияло на оригинал")
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key", 123)

	if !m.Exists("key") {
		t.Error("Exists(\"key\") = false, ожидалось true")
	}
	if m.Exists("absent") {
		t.Error("Exists(\"absent\") = true, ожидалось false")
	}
}

func TestGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 100)
	m.Add("zero", 0) // ноль — валидное значение!

	// Существующие ключи
	if v, ok := m.Get("a"); !ok || v != 100 {
		t.Errorf("Get(\"a\") = (%d, %v), ожидалось (100, true)", v, ok)
	}
	if v, ok := m.Get("zero"); !ok || v != 0 {
		t.Errorf("Get(\"zero\") = (%d, %v), ожидалось (0, true)", v, ok)
	}

	// Отсутствующий ключ
	if v, ok := m.Get("missing"); ok || v != 0 {
		t.Errorf("Get(\"missing\") = (%d, %v), ожидалось (0, false)", v, ok)
	}
}
