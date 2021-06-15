package main

import (
	"fmt"
	"math"
	"sync"
)

type printer interface {
	Print([]int)
}

func printArray(arr []int, p printer) {

	var wg sync.WaitGroup
	n := int(math.Ceil(float64(len(arr)) / 5))

	srcChan := make(chan int, 5)
	wg.Add(n)
	go func() {
		for _, v := range arr {
			srcChan <- v

		}
		close(srcChan)
	}()
	for i := 0; i < n; i++ {

		go func() {
			newArr := make([]int, 0, 2)
			for i := 0; i < 5; i++ {
				v, ok := <-srcChan
				if ok == false {
					break
				}
				newArr = append(newArr, v)
			}
			p.Print(newArr)
			wg.Done()
		}()
	}

	wg.Wait()

}

type myPrinter struct{}

func (myPrinter) Print(arr []int) {
	fmt.Println(arr)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	var m myPrinter

	printArray(arr, m)

}
