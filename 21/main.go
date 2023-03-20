package main

import "fmt"

type Target interface {
	Request() string
}

// Adaptee адаптируемый интерфейс
type Adaptee interface {
	SpecificRequest() string
}

type adapteeImpl struct{}

// SpecificRequest реализация адаптируемого интерфейса
func (a adapteeImpl) SpecificRequest() string {
	return "specific request"
}

// сам адаптер,имплементирующий интерфейс Target
type adapter struct {
	Adaptee
}

// Request работа с адаптируемым интерфейсом
func (a adapter) Request() string {
	return a.SpecificRequest()
}

func main() {
	adaptee := &adapteeImpl{}
	target := &adapter{Adaptee: adaptee}
	fmt.Println(target.Request())
}
