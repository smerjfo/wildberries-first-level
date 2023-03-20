package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr []int) {
	low, high := 0, len(arr)-1
	if low >= high {
		return
	}
	//сортируем массив относительно опорного элемента,принимаем индекс этого элемента
	p := partition(arr, low, high)
	//вызываем сортировку для правого и левого подмассива
	if p > low {
		quickSort(arr[:p])
	}
	if p < high {
		quickSort(arr[p+1:])
	}

}

// выбираем рандомный элемент и делаем его опорным
func choosePivot(arr []int, low int, high int) int {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Choose a random index within the subarray
	randIndex := rand.Intn(high-low+1) + low

	// Swap the random element with the first element to make it the pivot
	arr[randIndex], arr[high] = arr[high], arr[randIndex]

	// Return the index of the pivot element
	return high
}

func partition(arr []int, low int, high int) int {
	pivot := choosePivot(arr, low, high)
	//j указывает место,где все элементы слева-меньше опорного,а все элементы справа-больше или равны
	j := low - 1
	for i := low; i < high; i++ {
		//если элемент меньше чем опорный,мы меняем его местами с первым элементом,который больше или равен
		//если такие есть,если нет-просто сдвигаем указатель
		if arr[i] < arr[pivot] {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	//ставим опорный элемент на отсортированное относительно него место
	arr[j+1], arr[pivot] = arr[pivot], arr[j+1]
	return j + 1
}

func main() {
	arr := []int{1, 3, 1, 4, 5, 5, 2, 3, 5, 6, 9}
	quickSort(arr)
	fmt.Println(arr)
}
