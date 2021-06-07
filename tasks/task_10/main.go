package main

import "fmt"

func build(arr [][]int) {
	for _, v := range arr {
		fmt.Println(v)
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
	if arr[x][y] == 0 || arr[x][y] > n {

		arr[x][y] = n + 1
		n = arr[x][y]
	} else {
		return
	}

	nextPosition(arr, x, y+1, n)
	nextPosition(arr, x+1, y, n)
	nextPosition(arr, x, y-1, n)
	nextPosition(arr, x-1, y, n)

}
func wayBack(arr [][]int, x, y, n int) {
	if x < 0 || x >= len(arr) {
		return
	}
	if y < 0 || y >= len(arr[0]) {
		return
	}
	if arr[x][y] == n {
		arr[x][y] = -1
		n = n - 1
		if n == 1 {
			return
		}
	} else {
		return
	}

	wayBack(arr, x+1, y, n)
	wayBack(arr, x, y+1, n)
	wayBack(arr, x-1, y, n)
	wayBack(arr, x, y-1, n)

}
func theWay(arr [][]int) {
	var way [][]int
	for _, row := range arr {
		way = append(way, row)
		for j, cell := range row {
			if cell == -1 {
				row[j] = 0
			} else {
				row[j] = 1
			}
		}
	}
	build(way)

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
	build(matrix)
	n := 1
	nextPosition(matrix, x, y, n)
	x = len(matrix) - 1
	y = len(matrix[0]) - 1
	n = matrix[x][y]
	if matrix[x][y] > 1 {
		fmt.Println("Here is the Way")
		fmt.Println("It takes ", n, "steps")
		wayBack(matrix, x, y, n)
		theWay(matrix)
	} else {
		fmt.Println("No Way")
	}
}
