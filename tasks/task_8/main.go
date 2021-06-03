package main

import "fmt"

func fibonachi(n uint) uint {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonachi(n-1) + fibonachi(n-2)
	}

}
func main() {
	fmt.Println(fibonachi(0))
}
