package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

// New принимает 2 координаты,возвращает объект Point
func New(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Distance возвращает расстояние между 2 point'ами
func (p *Point) Distance(point *Point) float64 {
	dx := point.x - p.x
	dy := point.y - p.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p := New(0, 8)
	q := New(10, 4)
	fmt.Println(p.Distance(q))
}
