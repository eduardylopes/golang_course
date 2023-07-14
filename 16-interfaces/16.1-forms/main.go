package main

import (
	"fmt"
	"math"
)

type Form interface {
	area() float64
}

func showArea(form Form) {
	fmt.Printf("The area of this form is %0.2fmm²\n", form.area())
}

type Rectangle struct {
	width, height float64
}

func (rectangle Rectangle) area() float64 {
	return rectangle.height * rectangle.width
}

type Circle struct {
	radius float64
}

func (circle Circle) area() float64 {
	return math.Pow(math.Pi*circle.radius, 2)
}

func main() {
	rectangle := Rectangle{10, 15}
	showArea(rectangle)

	circle := Circle{6}
	showArea(circle)
}
