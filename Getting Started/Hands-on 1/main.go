package main

import "fmt"

type square struct {
	side float64
}

type circle struct {
	radius float64
	pi     float64
}

func (s square) area() {
	fmt.Println(s.side * s.side)
}

func (c circle) area() {
	fmt.Println(c.pi * (c.radius * c.radius))
}

type shape interface {
	area()
}

func areaPrint(s shape) {
	s.area()
}

func main() {
	s1 := square{
		2.0,
	}

	c1 := circle{
		10.0,
		3.1415,
	}

	areaPrint(s1)
	areaPrint(c1)
}
