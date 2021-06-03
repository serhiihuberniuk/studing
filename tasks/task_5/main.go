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

func rectanglePer(r rectangle) float64 {
	return r.a*2 + r.b*2
}

func squarePer(s square) float64 {
	return s.a * 4
}

func trianglePer(t triangle) float64 {
	return t.a + t.b + t.c
}

func rectangleSq(r rectangle) float64 {
	return r.a * r.b
}

func squareSq(s square) float64 {
	return s.a * s.a
}

func triangleSq(t triangle) float64 {
	halfP := trianglePer(t) / 2
	return math.Sqrt(halfP * (halfP - t.a) * (halfP - t.b) * (halfP - t.c))
}

/* methods
func(r rectangle) peri()float64{
	return  r.a*2 + r.b*2
}
func(s square) peri()float64{
	return s.a *4
}
func (t triangle) peri()float64{
	return  t.a+t.b+t.c
}
func (r rectangle) square()float64{
	return  r.a * r.b
}
func(t triangle) square()float64{
	halfP := trianglePer(t)/2
	return math.Sqrt(halfP*(halfP-t.a)*(halfP - t.b) * (halfP - t.c))
}
func(s square) square()float64{
	return  s.a*s.a
}
*/

func totalSquare(r rectangle, t triangle, s square) float64 {
	return rectangleSq(r) + triangleSq(t) + squareSq(s)
}
func main() {
	r := rectangle{10, 15}
	s := square{10}
	t := triangle{10, 15, 20}

	fmt.Println("Rectangle P =", rectanglePer(r), "S =", rectangleSq(r))
	fmt.Println("Square P =", squarePer(s), "S =", squareSq(s))
	fmt.Println("Triangle P =", trianglePer(t), "S =", triangleSq(t))

	fmt.Println("total square: ", totalSquare(r, t, s))
}
