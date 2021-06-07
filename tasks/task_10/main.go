package main

import (
	"errors"
	"fmt"
)

func printArr(arr [][]int, wayIsFound bool) {
	for _, row := range arr {
		fmt.Println()
		for _, cell := range row {
			if wayIsFound {
				if cell != -1 {
					fmt.Print(1, " ")
				} else {
					fmt.Print(0, " ")
				}
			} else {
				fmt.Print(cell, " ")
			}
		}
	}
	fmt.Println()
}

func nextPosition(arr [][]int, x, y, n int) {
	if x < 0 || x >= len(arr) {
		return
	}
	if y < 0 || y >= len(arr[0]) {
		return
	}
	if arr[x][y] == 1 {
		return
	}
	if arr[x][y] < n && arr[x][y] != 0 {
		return
	}
	arr[x][y] = n + 1
	n = arr[x][y]

	nextPosition(arr, x, y+1, n)
	nextPosition(arr, x+1, y, n)
	nextPosition(arr, x, y-1, n)
	nextPosition(arr, x-1, y, n)

}
func backWay(arr [][]int) ([][]int, error) {
	x, y := len(arr)-1, len(arr[0])-1
	n := arr[x][y]
	for {
		arr[x][y] = -1
		if x == 0 && y == 0 {
			return arr, nil
		}
		if y+1 < len(arr[0]) && arr[x][y+1] == n-1 {
			y++
			n--
		} else if y-1 >= 0 && arr[x][y-1] == n-1 {
			y--
			n--
		} else if x-1 >= 0 && arr[x-1][y] == n-1 {
			x--
			n--
		} else if x+1 < len(arr) && arr[x+1][y] == n-1 {
			y++
			n--
		} else {
			var err error = errors.New("invalid array")
			return nil, err
		}

	}

}

func main() {
	var x, y int
	matrix := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 1, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	printArr(matrix, false)
	n := 1
	nextPosition(matrix, x, y, n)

	if matrix[x][y] > 1 {
		matrix, err := backWay(matrix)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Here is the Way")
			printArr(matrix, true)
		}
	} else {
		fmt.Println("No Way")
	}
}
