package chapter4

import (
	"fmt"
	"math"
)

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

type geometry interface {
	area() float64
	perim() float64
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func Interface() {
	r := Rect{3.5, 6.3}
	c := Circle{7.4}

	measure(&r)
	measure(&c)
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perim() float64 {
	return 2 * r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
