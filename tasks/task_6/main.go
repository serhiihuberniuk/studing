package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	a, b float64
}

type square struct {
	a float64
}
type triangle struct {
	a, b, c float64
}

func (r rectangle) peri() float64 {
	return r.a*2 + r.b*2
}
func (s square) peri() float64 {
	return s.a * 4
}
func (t triangle) peri() float64 {
	return t.a + t.b + t.c
}
func (r rectangle) square() float64 {
	return r.a * r.b
}
func (t triangle) square() float64 {
	halfP := t.peri() / 2
	return math.Sqrt(halfP * (halfP - t.a) * (halfP - t.b) * (halfP - t.c))
}
func (s square) square() float64 {
	return s.a * s.a
}

func totalSquare(r rectangle, t triangle, s square) float64 {
	return r.square() + s.square() + t.square()
}
func main() {
	r := rectangle{10, 15}
	s := square{10}
	t := triangle{10, 15, 20}

	fmt.Println("Rectangle P =", r.peri(), "S =", r.square())
	fmt.Println("Square P =", s.peri(), "S =", s.square())
	fmt.Println("Triangle P =", t.peri(), "S =", t.square())

	fmt.Println("total square: ", totalSquare(r, t, s))
}
