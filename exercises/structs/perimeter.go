package main

import "math"

// Shape generic method
type Shape interface {
	Area() float64
}

// Rectangle ...
type Rectangle struct {
	Width  float64
	Height float64
}

// Area ....
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle ...
type Circle struct {
	Radius float64
}

// Area ....
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter ...
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func main() {

}
