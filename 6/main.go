package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	goroutineStopByChannel()
	goroutineStopByContextWithTimeout()
	goroutineStopByContextWithCancel()
	goroutineStopByBooleanFlag()
}

func goroutineStopByChannel() {
	ch := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch:
				fmt.Println("Завершение горутины при помощи канала.")
				return
			}
		}
	}()
	ch <- struct{}{}
	wg.Wait()
}

func goroutineStopByContextWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Завершение горутины при помощи контекста с таймаутом.")
				return
			}
		}
	}()
	wg.Wait()
}
func goroutineStopByContextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Завершение горутины при помощи контекста.")
				return
			}
		}
	}()
	cancel()
	wg.Wait()
}

func goroutineStopByBooleanFlag() {
	var stopped bool
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for !stopped {
			// do some work
		}
		fmt.Println("Завершение горутины при помощи флага.")
		wg.Done()
	}()

	// To stop the goroutine, set the stopped variable to true:
	stopped = true
	wg.Wait()
}
