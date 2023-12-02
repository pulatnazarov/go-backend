package main

import (
	"fmt"
)

// Animal interface with method and printSound func

func main() {
	s := Square{
		Length: 5,
	}

	c := Circle{
		Radius: 3,
	}

	printShapeSurface(s)
	printShapePerimeter(s)

	printShapeSurface(c)
	printShapePerimeter(c)
}

type Shape interface {
	Surface()
	Perimeter()
}

func printShapeSurface(s Shape) {
	s.Surface()
}

func printShapePerimeter(s Shape) {
	s.Perimeter()
}

type Square struct {
	Length int
}

func (s Square) Surface() {
	fmt.Println("surface of square: ", s.Length*s.Length)
}

func (s Square) Perimeter() {
	fmt.Println("perimeter of square: ", 4*s.Length)
}

type Circle struct {
	Radius int
}

func (c Circle) Surface() {
	fmt.Println("surface of circle: ", 3*c.Radius)
}

func (c Circle) Perimeter() {
	fmt.Println("surface of circle: ", 2*3*c.Radius)
}
