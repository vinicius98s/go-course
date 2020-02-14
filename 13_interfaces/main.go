package main

import (
	"fmt"
	"math"
)

// Defining structs Circle and Retangle

// Circle is nice
type Circle struct {
	x, y, radius float64
}

// Rectangle is nice as well
type Rectangle struct {
	width, height float64
}

// Interface Shape: wraps Circle and Retangle (needs a area function)

// Shape Define interface
type Shape interface {
	area() float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}
