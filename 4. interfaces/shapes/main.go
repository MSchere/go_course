package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	side float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	square := square{5.2}
	triangle := triangle{5.2, 2.5}
	printArea(square)
	printArea(triangle)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (s square) getArea() float64 {
	return s.side * s.side
}
