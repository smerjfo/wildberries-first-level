package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	//вводим количество секунд
	var n int
	for {
		fmt.Print("Введите количество секунд исполнения: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Println("Ошибка ввода: ", err)
		} else {
			fmt.Println("Вводите текст:")
			break
		}
	}

	ch := make(chan string)
	wg := sync.WaitGroup{}
	//создаем контекст с таймаутом в n секунд
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()

	wg.Add(1)

	go func(chOut <-chan string) {
		defer wg.Done()
		//читаем данные из канала и выводим их в консоль,ожидая,пока пройдет n секунд
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Завершаем работу...\n")
				return
			case a := <-chOut:
				fmt.Print(a)
			}
		}
	}(ch)
	//читаем данные из stdin и пишем их в канал
	go func(chIn chan<- string) {
		for {
			input := Scan()
			chIn <- input
		}
	}(ch)
	<-ctx.Done()

	wg.Wait()
	fmt.Println("Работа успешно завершена.")
}

func Scan() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	return str
}
