package main

import (
	"fmt"
	"sync"
)

// реализация конкурентной записи в map
func main() {
	writes := 1000
	storage := make(map[int]int)
	//wg для ожидания окончания исполнения всех горутин
	wg := sync.WaitGroup{}
	wg.Add(writes)
	//синхронизация записи в map с помощью mutex
	mu := sync.Mutex{}
	for i := 0; i < writes; i++ {
		//передача в анонимную функцию переменной по значению
		i := i
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			storage[i] = i
		}()
	}
	wg.Wait()
	fmt.Println(storage)
}
