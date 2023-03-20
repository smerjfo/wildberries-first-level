package main

import "fmt"

func main() {
	arr := []int{0}
	fmt.Println(deletion(arr, 1))
}

func deletion(arr []int, index int) []int {
	//проверка что такой индекс есть в массиве и что массив не пустой
	if index < 0 || len(arr) == 0 || index >= len(arr) {
		return arr
	}
	//удаляем элемент
	newArray := append(arr[:index], arr[index+1:]...)
	return newArray
}
