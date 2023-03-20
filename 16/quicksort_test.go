package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	var tests = [][]int{
		{1, 3, 1, 4, 5, 5, 2, 3, 5, 6, 9},
		{},
		{1},
		{2, 2, 1},
	}
	for _, test := range tests {
		quickSort(test)
		checkArray(test, t)
		t.Log(test)
	}
	for i := 0; i < 100; i++ {
		arr := generateRandomSlice()
		quickSort(arr)
		checkArray(arr, t)
		t.Log(arr)
	}

}

func checkArray(arr []int, t *testing.T) {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			t.Errorf("quickSort failed: %v is not sorted in ascending order", arr)
			break
		}
	}
}

func generateRandomSlice() []int {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random length for the slice
	length := rand.Intn(100) + 1 // Minimum length of 1, maximum length of 10

	// Initialize the slice with the random length
	slice := make([]int, length)

	// Fill the slice with random elements
	for i := 0; i < length; i++ {
		slice[i] = rand.Intn(100) // Elements will be between 0 and 99
	}

	return slice
}
