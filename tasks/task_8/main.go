package main

import "fmt"

func fibonachi(n uint) (f uint) {
	if n == 0 {
		f = 0
	} else if n == 1 {
		f = 1
	} else {
		fiboArray := []uint{0, 1}
		for i := 2; uint(i) <= n; i++ {
			f = fiboArray[i-1] + fiboArray[i-2]
			fiboArray = append(fiboArray, f)
		}
	}
	return f

}
func main() {
	fmt.Println(fibonachi(10))
}
