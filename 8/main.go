package main

import (
	"fmt"
)

func main() {
	var val int64
	var i uint

	//вводим число и индекс
	fmt.Print("Введите число и индекс: ")
	_, err := fmt.Scan(&val, &i)
	if err != nil {
		fmt.Println(err)
		return
	}
	//создаем маску для i'ого бита,где i бит-1
	oneBitMask := int64(1) << i
	fmt.Println(oneBitMask)
	//обратная маска
	zeroBitMask := ^oneBitMask
	//если i'ый бит в переданном числе 1,то мы его обнуляем
	//иначе мы его устанавливаем в 1
	if oneBitMask&val > 0 {
		fmt.Printf("У числа %d бит №%d не нулевой\n", val, i)
		val = zeroBitMask & val
	} else {
		fmt.Printf("У числа %d бит №%d нулевой\n", val, i)
		val = oneBitMask | val
	}
	fmt.Println("Результат: ", val)

}
