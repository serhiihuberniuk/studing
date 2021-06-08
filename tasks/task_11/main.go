package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func addZeros(s string, ch1 chan string) {
	s = s + "000"
	ch1 <- s
}
func convertToInt(ch1 chan string, ch2 chan int) {
	s := <-ch1
	i, err := strconv.Atoi(s)
	if err != nil {
		err = errors.New("wrong string")
	}
	ch2 <- i
}

func convertToString(ch2 chan int) {
	i := <-ch2
	s := strconv.Itoa(i)
	fmt.Println(s)
}

func main() {
	text := "90 23 239 123 34 4452 23 123 34 12 34"
	var stringArray []string
	stringArray = strings.Fields(text)
	for _, v := range stringArray {
		var ch1 = make(chan string)
		var ch2 = make(chan int)
		go addZeros(v, ch1)
		go convertToInt(ch1, ch2)
		go convertToString(ch2)
	}
	var input string   // stole this decision from golang-book
	fmt.Scanln(&input) // do i had to use mutex?

}
