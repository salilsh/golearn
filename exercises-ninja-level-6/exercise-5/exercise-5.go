package main

import (
	"fmt"
)

type Circle struct {
	radius float64
}

type Rect struct {
	length float64
	width  float64
}

func (c Circle) area() float64 {
	circleArea := 3.14 * c.radius * c.radius
	return circleArea
}

func (r Rect) area() float64 {
	rectArea := r.length * r.width
	return rectArea
}

type polyMorphicArea interface {
	area() float64
}

func calculateArea(p polyMorphicArea) float64 {
	return p.area()
}

func main() {
	rec := Rect{
		length: 5,
		width:  6,
	}

	circ := Circle{
		radius: 5,
	}

	fmt.Println(calculateArea(rec))
	fmt.Println(calculateArea(circ))
}
