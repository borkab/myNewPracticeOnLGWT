package peri

import "math"

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Height   float64
	Baseline float64
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

func (rec Rectangle) Perimeter() float64 {
	return 2 * (rec.Height + rec.Width)
}

func (cir Circle) Perimeter() float64 {
	return 2 * (cir.Radius * math.Pi)
}

func (rec Rectangle) Area() float64 {
	return rec.Height * rec.Width
}

func (cir Circle) Area() float64 {
	return cir.Radius * cir.Radius * math.Pi
}

func (t Triangle) Perimeter() float64 {
	return t.Baseline + t.Height + 17
}

func (t Triangle) Area() float64 {
	return (t.Baseline * t.Height) * 0.5
}
