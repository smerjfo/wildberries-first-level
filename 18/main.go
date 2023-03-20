package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	//значение счетчика и мьютекс связанный с ним
	val int
	mu  sync.RWMutex
}

func New() *Counter {
	return &Counter{
		val: 0,
		mu:  sync.RWMutex{},
	}
}

// Increment concurrency-safe increment
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++

}

// GetValue concurrency-safe reading of value counter
func (c *Counter) GetValue() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.val
}

func main() {
	c := New()
	wg := sync.WaitGroup{}
	wg.Add(5)
	//пишем в счетчик из 5 горутин
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				c.Increment()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.GetValue())
}
