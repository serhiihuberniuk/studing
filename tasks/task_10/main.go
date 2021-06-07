package main

import "fmt"

func printArr(arr [][]int) {
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

func backWay(arr [][]int, x, y, n int) {
	n = arr[x][y]
	arr[x][y] = -1
	for {

		for {

			a := x
			b := y + 1
			if (a >= 0 && a < len(arr)) && (b >= 0 && b < len(arr[0])) {
				if arr[a][b] == n-1 {
					x = a
					y = b
					arr[x][y] = -1
					n = n - 1
					break
				}
			}
			a = x + 1
			b = y
			if (a >= 0 && a < len(arr)) && (b >= 0 && b < len(arr[0])) {
				if arr[a][b] == n-1 {
					x = a
					y = b
					arr[x][y] = -1
					n = n - 1
					break
				}
			}
			a = x
			b = y - 1
			if (a >= 0 && a < len(arr)) && (b >= 0 && b < len(arr[0])) {
				if arr[a][b] == n-1 {
					x = a
					y = b
					arr[x][y] = -1
					n = n - 1
					break
				}
			}
			a = x - 1
			b = y
			if (a >= 0 && a < len(arr)) && (b >= 0 && b < len(arr[0])) {
				if arr[a][b] == n-1 {
					x = a
					y = b
					arr[x][y] = -1
					n = n - 1
					break
				}
			}
		}
		if x == 0 && y == 0 {
			break
		}
	}
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
	printArr(way)
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
	printArr(matrix)
	n := 1
	nextPosition(matrix, x, y, n)
	x = len(matrix) - 1
	y = len(matrix[0]) - 1
	n = matrix[x][y]
	if matrix[x][y] > 1 {
		fmt.Println("Here is the Way")
		fmt.Println("It takes ", n, "steps")
		backWay(matrix, x, y, n)
	} else {
		fmt.Println("No Way")
	}
}
