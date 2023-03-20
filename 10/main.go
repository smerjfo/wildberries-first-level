package main

import "fmt"

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	values := make(map[int][]float64)
	for _, v := range arr {
		//отбрасываем дробную часть,затем убираем единицы
		iv := int(v) / 10 * 10
		//создаем массив по ключу в мапе
		if _, ok := values[iv]; ok != true {
			values[iv] = make([]float64, 0)
		}
		//добавляем значение в нужный массив
		array, _ := values[iv]
		values[iv] = append(array, v)
	}
	fmt.Println(values)
}
