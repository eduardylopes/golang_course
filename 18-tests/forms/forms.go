package forms

import (
	"fmt"
	"math"
)

type Form interface {
	Area() float64
}

func ShowArea(form Form) {
	fmt.Printf("The area of this form is %0.2fmm²\n", form.Area())
}

type Rectangle struct {
	Width, Height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Height * rectangle.Width
}

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.Radius, 2)
}
