package main

import "fmt"

const (
	PI = 3.14
)

type Circle struct {
	radius float64
}

func NewCircle(radius float64) Circle {
	return Circle{
		radius: radius,
	}
}

func (c Circle) calculateCircumference() {
	circumference := 2 * PI * c.radius
	fmt.Println(circumference)
}

func main() {
	myCircle := NewCircle(1.5)

	myCircle.calculateCircumference()
}
