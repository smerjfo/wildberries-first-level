package main

func main() {
	firstArr := []int{1, 2, 3, 4, 12, 19}
	secondArr := []int{9, 10, 73, 48, 12, 1}
	intersectionTwoArrays(firstArr, secondArr)
}

func intersectionTwoArrays(firstArr, secondArr []int) []int {

	//мапа будет хранить ключи из первого массива
	intersectionMap := make(map[int]struct{})
	result := make([]int, 0)

	//запись в мапу значений из первого множества
	for _, i2 := range firstArr {
		intersectionMap[i2] = struct{}{}
	}

	//проверка есть ли в мапе значение из второго множества
	for _, i2 := range secondArr {
		//если есть,добавляем их в результат
		if _, ok := intersectionMap[i2]; ok == true {
			result = append(result, i2)
		}
	}
	return result
}
