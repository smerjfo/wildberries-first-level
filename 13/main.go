package main

import "fmt"

func main() {
	i, y := 1, 12
	i, y = y, i
	fmt.Println("I :", i, " Y:", y)
}
