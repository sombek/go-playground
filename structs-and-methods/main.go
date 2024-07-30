package structs_and_methods

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}
type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Shape interface {
	Area() float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Circle) Area() float64 {
	return r.Radius * r.Radius * math.Pi
}
