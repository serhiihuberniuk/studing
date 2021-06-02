package main

import (
	"fmt"
)

func main() {
	text := "abc"
	fmt.Println(textFromTheEnd(text))

}

func textFromTheEnd(text string) string {
	stringByRunes := []rune(text)
	for i := 0; i < len(stringByRunes)/2; i++ {
		stringByRunes[i], stringByRunes[len(stringByRunes)-1-i] = stringByRunes[len(stringByRunes)-1-i], stringByRunes[i]

	}
	return string(stringByRunes)
}
