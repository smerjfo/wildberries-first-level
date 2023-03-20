package main

import "fmt"

// Search поиск индекса элемента target в массиве
func Search(nums []int, target int) int {
	var half int
	var index int

	if len(nums)%2 != 0 {
		half = len(nums) / 2
	} else {
		half = (len(nums) - 1) / 2
	}
	//индекс отслеживает положение во всем массиве,half это индекс середины в текущем слайсе
	index = half

	for len(nums) != 0 {
		if nums[half] == target {
			return index
		}
		//если серединный элемент больше искомого,работаем с левой частью
		if nums[half] > target {
			nums = nums[:half]
			//устанавливливаем указатель на новый серединный элемент,смещаем индекс
			if len(nums)%2 == 0 {
				half = (len(nums) - 1) / 2
				index -= half + 2
			} else {
				half = len(nums) / 2
				index -= half + 1
			}
			//иначе работаем с правой частью
		} else {
			nums = nums[half+1:]
			if len(nums)%2 == 0 {
				half = (len(nums) - 1) / 2
				index += half + 1
			} else {
				half = len(nums) / 2
				index += half + 1
			}
		}
	}
	//элемент не найден
	return -1
}

// поиск через 2 указателя
func searchTwoPointers(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		//берем центральный элемент и сравниваем с искомым
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	//элемент не найден
	return -1
}

func main() {
	fmt.Println(Search([]int{1, 2, 3, 419, 5000, 100000}, 419))
}
