package main

import (
	"fmt"
	"studing/tasks/pkg/geometry"
)

type house struct {
	roof   geometry.Triangle
	ground geometry.Rectangle
	wall   geometry.Square
}

func main() {
	var home house
	home.wall.A = 13 // not sure its the best way to value the fields
	home.roof.A = 8
	home.roof.B = 8
	home.roof.C = 6

	squareToPaint := home.wall.Square()*4 + home.roof.Square()*2
	fmt.Println(squareToPaint)

}
