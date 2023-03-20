package main

import "fmt"

func main() {
	var val interface{}
	val = make(chan interface{})
	checkType(val)
	val = 1
	checkType(val)
	val = "123"
	checkType(val)
	val = false
	checkType(val)
}

func checkType(i interface{}) {
	//динамическое определение типа переданного параметра
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %v\n", v)
	case string:
		fmt.Printf("String: %v\n", v)
	case bool:
		fmt.Printf("Bool: %v\n", v)
	case chan interface{}:
		fmt.Printf("This is channel\n")
	default:
		fmt.Printf("Unknown\n")
	}

}
