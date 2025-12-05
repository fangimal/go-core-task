package main

import (
	"sync/atomic"
)

type CustomWaitGroup struct {
	counter int64
	done    chan struct{}
}

func (wg *CustomWaitGroup) Add(delta int) {
	newCount := atomic.AddInt64(&wg.counter, int64(delta))

	if newCount < 0 {
		panic("sync: negative WaitGroup counter")
	}

	if newCount == 0 && wg.done != nil {
		close(wg.done)
		wg.done = nil
	}
}

func (wg *CustomWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *CustomWaitGroup) Wait() {
	if atomic.LoadInt64(&wg.counter) == 0 {
		return
	}

	if wg.done == nil {
		wg.done = make(chan struct{})
	}

	<-wg.done
}
