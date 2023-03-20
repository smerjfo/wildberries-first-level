package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	chX := make(chan int)
	chXSquare := make(chan int)
	//читаем из массива данные и пишем в канал
	go func() {
		for i, _ := range arr {
			chX <- arr[i]
		}
		close(chX)
	}()
	//возводим данные переданные в канал в квадрат и передаем дальше
	go func() {
		for v := range chX {
			chXSquare <- v * v
		}
		close(chXSquare)
	}()
	wg := sync.WaitGroup{}
	wg.Add(1)
	//читаем результаты и пишем в stdout
	go func() {
		defer wg.Done()
		for v := range chXSquare {
			fmt.Println(v)
		}
	}()
	wg.Wait()
}
