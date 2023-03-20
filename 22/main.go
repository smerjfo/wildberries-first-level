package main

import (
	"fmt"
	"math/big"
)

type Number struct {
	big.Int
}

// New конструктор принимает число в виде строки,возвращает экземпляр Number
func New(num string) *Number {
	number := new(big.Int)
	number.SetString(num, 10)
	return &Number{Int: *number}
}

// Add принимает аргументы x,y,складывает их и помещает результат в n. Возвращает n.
func (n *Number) Add(x *Number, y *Number) *Number {
	n.Int.Add(&x.Int, &y.Int)
	return n
}

// Mul принимает аргументы x,y,умножает их и помещает результат в n. Возвращает n.
func (n *Number) Mul(x *Number, y *Number) *Number {
	n.Int.Mul(&x.Int, &y.Int)
	return n
}

// Div принимает аргументы x,y,делит x/y и помещает результат в n. Возвращает n.
func (n *Number) Div(x, y *Number) *Number {
	n.Int.Div(&x.Int, &y.Int)
	return n
}

// Sub принимает аргументы x,y,вычитает x-y и помещает результат в n. Возвращает n.
func (n *Number) Sub(x, y *Number) *Number {
	n.Int.Sub(&x.Int, &y.Int)
	return n
}

func main() {
	a := New("10000000000000000000000000000000000")
	b := New("15000")
	c := New("0")
	c.Sub(a, b)
	fmt.Println(c)
	c.Add(c, b)
	fmt.Println(c)
	c.Sub(a, b)
	c.Mul(a, b)
	fmt.Println(c)
}
