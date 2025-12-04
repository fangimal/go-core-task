package main

import (
	"testing"
	"time"
)

func TestCustomWaitGroup(t *testing.T) {
	var wg CustomWaitGroup

	wg.Add(2)

	go func() {
		time.Sleep(10 * time.Millisecond)
		wg.Done()
	}()

	go func() {
		time.Sleep(20 * time.Millisecond)
		wg.Done()
	}()

	start := time.Now()
	wg.Wait()
	duration := time.Since(start)

	if duration < 20*time.Millisecond {
		t.Errorf("Wait завершился слишком рано: %v", duration)
	}
}

func TestWaitGroupZero(t *testing.T) {
	var wg CustomWaitGroup
	wg.Wait()
}

func TestWaitGroupNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("ожидалась паника при отрицательном счётчике")
		}
	}()
	var wg CustomWaitGroup
	wg.Done()
}
