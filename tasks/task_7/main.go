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
	home := house{
		roof:   geometry.Triangle{8, 8, 6},
		ground: geometry.Rectangle{5, 15},
		wall:   geometry.Square{13},
	}

	squareToPaint := home.wall.Square()*4 + home.roof.Square()*2
	fmt.Println(squareToPaint)

}
