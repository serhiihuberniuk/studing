package main

import (
	"fmt"
	"strconv"
)

func addZeros(s string, out chan string) {
	s = s + "000"
	out <- s
}
func convertToInt(in chan string, out chan int) {
	s := <-in
	i, err := strconv.Atoi(s)
	if err != nil {
		out <- 1
		return
	} else {
		out <- i
	}
}

func convertToString(in chan int, out chan string) {

	i := <-in
	if i == 1 {
		out <- ""
		return
	} else {
		s := strconv.Itoa(i)
		out <- s
	}
}

func main() {
	stringArray := []string{"90", "23", "hh", "123", "34", "4452", "23", "123", "000", "12", "34"}
	var ch1 = make(chan string)
	var ch2 = make(chan int)
	var ch3 = make(chan string)
	for _, v := range stringArray {
		go addZeros(v, ch1)
		go convertToInt(ch1, ch2)
		go convertToString(ch2, ch3)
		s := <-ch3
		if s != "" {
			fmt.Println(s)
		}
	}

}
