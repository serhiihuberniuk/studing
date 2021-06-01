package main

import "fmt"

func task1(slice []int) {
	for _, v := range slice {
		if v%2 == 0 && v != 0 {
			fmt.Println(v)
		}
	}
}
func main() {
	slice := []int{1, 2, 17, 3, 45, 3, 56, 78, -9, -4, 0}
	task1(slice)
}
