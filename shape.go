package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	ss := square{sideLength: 2}
	ts := triangle{height: 2, base: 2}

	fmt.Println("The area of your square is ", getArea(ss))
	fmt.Println("The area of your triangle is ", getArea(ts))
}

func getArea(s shape) float64 {
	return s.getArea()
}

func (t triangle) getArea() float64 {
	area := (t.height / 2) * t.height
	return area
}

func (s square) getArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}
