package structs

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Perimeter() (perimeter float64) {
	return (r.Height + r.Width) * 2
}

func (r Circle) Perimeter() (perimeter float64) {
	return 2 * r.Radius * math.Pi
}

func (r Rectangle) Area() (area float64) {
	return r.Height * r.Width
}

func (r Circle) Area() (area float64) {
	return math.Pow(r.Radius, 2) * math.Pi
}
