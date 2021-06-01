package main

import "fmt"

func task2(oldSlice []int) []int {
	var newSlice []int
	for _, v := range oldSlice {
		v *= 2
		newSlice = append(newSlice, v)
	}
	return newSlice
}

func main() {
	theSlice := []int{1, 15, 18, 23, 0, 156, -89, -1}
	fmt.Println(task2(theSlice))
}
