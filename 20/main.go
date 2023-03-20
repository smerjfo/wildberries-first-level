package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		str := Scan()
		fmt.Println(reverseOrderOfWords(str))
	}
}

// изменяем порядок слов на обратный
func reverseOrderOfWords(entry string) string {
	//получили массив слов,разделенных пробелами
	words := strings.Split(entry, " ")
	//изменили порядок на обратный
	for i, l := 0, len(words)-1; i < len(words)/2; i++ {
		words[i], words[l-i] = words[l-i], words[i]
	}
	//используя builder конкатенируем массив слов в строку
	result := strings.Builder{}
	for _, word := range words {
		result.WriteString(word)
		result.WriteByte(' ')
	}
	return result.String()
}

// Scan функция считывания строки из os.Stdin
func Scan() string {
	fmt.Print("Введите строку: ")
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	return strings.Trim(str, "\n")
}
