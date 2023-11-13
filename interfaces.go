package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return 4 * s.length
}

type shape interface {
	area() float64
	circumf() float64
}

func printShapeInfo(s shape) {
	fmt.Printf("area: %0.2f\n", s.area())
	fmt.Printf("circumference: %0.2f\n", s.circumf())
}

func main() {
	shapes := []shape{ circle{10.0}, square{10.0} }
	for _, v := range shapes {
		printShapeInfo(v)
	}
}
