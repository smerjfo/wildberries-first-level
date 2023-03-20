package main

import (
	"fmt"
	"sync"
)

var arr = [5]int{2, 4, 6, 8, 10}

func main() {

	wg := sync.WaitGroup{}
	//Запускаем для каждого числа из массива arr собственную goroutine
	for _, i2 := range arr {
		wg.Add(1)
		//i2 передаем по значению
		go func(i2 int) {
			fmt.Println(i2 * i2)
			wg.Done()
		}(i2)
	}
	//Ждем окончания исполнения goroutine
	wg.Wait()
}
