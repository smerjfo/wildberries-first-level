package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "aBbcDrqwkejl"
	fmt.Println(isAllUnique(str))
}

// проверка что все символы в строку уникальны,регистронезависимый результат
func isAllUnique(str string) bool {
	str = strings.ToLower(str)
	//проверяем есть ли уже в множестве такой символ.
	//если есть сразу возвращаем false
	set := make(map[rune]struct{})
	runes := []rune(str)
	for _, r := range runes {
		if _, ok := set[r]; ok == true {
			return false
		}
		set[r] = struct{}{}
	}
	return true
}
