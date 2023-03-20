package main

import "fmt"

func main() {
	var arr = []int{2, 4, 6, 8, 10}
	fmt.Println(concurrentSum(arr))
}

func concurrentSum(arr []int) int {
	var sum int
	ch := make(chan int)
	//запускаем goroutine для чисел из массива и пишем результат вычисления в канал
	for _, val := range arr {
		go func(i int) {
			ch <- i * i
		}(val)
	}
	//пишем результат вычислений
	for range arr {
		sum += <-ch
	}
	return sum
}
