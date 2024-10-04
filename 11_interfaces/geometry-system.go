package main

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (rect rectangle) area() float64 {
	return rect.width * rect.height
}

func (rect rectangle) perimeter() float64 {
	return 2 * (rect.width + rect.height)
}

func measure(s shape) {
	s.area()
	s.perimeter()
}
