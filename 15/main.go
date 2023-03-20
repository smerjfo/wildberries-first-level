package main

import (
	"fmt"
	"math/rand"
	"time"
)

func someFunc() (string, error) {
	v := []rune(createHugeString(1 << 10))
	//ожидаем строку не менее 100 символов
	if len(v) >= 100 {
		return string(v[:100]), nil
	} else {
		return "", fmt.Errorf("Строка меньше 100 символов")
	}

}

// создаем большую строку
func createHugeString(i int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())

	size := i
	runes := make([]rune, size)
	//рандомно выбираем букву и записываем пока строка не будет нужного размера
	for i := 0; i < size; i++ {
		runes[i] = rune(letters[rand.Intn(len(letters))])
	}

	randomString := string(runes)
	return randomString
}

func main() {
	str, err := someFunc()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
}
