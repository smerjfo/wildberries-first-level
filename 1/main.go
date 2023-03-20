package main

import "fmt"

type Human struct {
}

func (h Human) hello() {
	fmt.Println("I'm a human")
}

type Action struct {
	Human
}

func main() {
	a := Action{Human: Human{}}
	a.hello()
}
