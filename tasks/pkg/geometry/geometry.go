package geometry

import (
	"math"
)

type Rectangle struct {
	A, B float64
}

func (r Rectangle) Peri() float64 {
	return r.A*2 + r.B*2
}
func (r Rectangle) Square() float64 {
	return r.A * r.B
}

type Square struct {
	A float64
}

func (s Square) Peri() float64 {
	return s.A * 4
}
func (s Square) Square() float64 {
	return s.A * s.A
}

type Triangle struct {
	A, B, C float64
}

func (t Triangle) Peri() float64 {
	return t.A + t.B + t.C
}

func (t Triangle) Square() float64 {
	halfP := t.Peri() / 2
	return math.Sqrt(halfP * (halfP - t.A) * (halfP - t.B) * (halfP - t.C))
}

func TotalSquare(r Rectangle, t Triangle, s Square) float64 {
	return r.Square() + s.Square() + t.Square()
}
