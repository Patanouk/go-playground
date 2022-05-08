package _struct

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Triangle struct {
	x,y float64
}

func (t Triangle) Area() float64 {
	return (t.x * t.y) / 2
}

func (t Triangle) Perimeter() float64 {
	return t.x + t.y + math.Sqrt(math.Pow(t.x, 2) + math.Pow(t.y, 2))
}
