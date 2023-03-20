package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	//выбираем количество воркеров
	var n int
	for {
		fmt.Print("Введите количество воркеров: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Println("Ошибка ввода: ", err)
		} else {
			break
		}
	}
	//канал передачи входящих данных
	ch := make(chan string)
	//канал передачи системных сигналов
	sigCh := make(chan os.Signal, 1)
	//канал через который останавливаем работу worker'ов
	stopChan := make(chan struct{})
	//если встречаем CTRL+C-передаем его в канал
	signal.Notify(sigCh, os.Interrupt)

	//wg нужна для graceful shutdown
	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Начало работы %d ...\n", id)
			//ждем данные или сигнал о завершении работы
			for {
				select {
				case <-stopChan:
					fmt.Printf("Завершаем работу %d ...\n", id)
					return
				case a := <-ch:
					fmt.Print(a)
				}
			}
		}(i)
	}
	go func() {
		for {
			input := Scan()
			ch <- input
		}
	}()
	//Ждем передачи CTRL+C
	<-sigCh
	fmt.Println("\nПолучен сигнал завершения работы,ожидаем завершения воркеров...")
	//уведомляем горутины о завершении программы
	close(stopChan)
	//ждем завершения всех горутин
	wg.Wait()
	fmt.Println("Все воркеры остановлены.")
}

// Scan функция считывания строки из os.Stdin
func Scan() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	return str
}
