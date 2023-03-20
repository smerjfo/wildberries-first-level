package main

import "fmt"

func main() {
	for {
		fmt.Print("Введите строку: ")
		var entry string
		fmt.Scan(&entry)
		fmt.Println("Результат:", reverse(entry))
	}

}

func reverse(entry string) string {
	//кастим строку к рунам,для корректного отображения 2-байтных символов
	runes := []rune(entry)
	length := len(runes) - 1
	//переворачиваем руны
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[length-i] = runes[length-i], runes[i]
	}
	//возвращаем строку
	return string(runes)
}
