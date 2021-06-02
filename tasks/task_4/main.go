package main

import (
	"fmt"
)

func main() {
	text := " a b c d e f g "
	fmt.Println(textFromTheEnd(text))

}

func textFromTheEnd(text string) string {
	stringByRunes := []rune(text)
	for i := 0; i < len(stringByRunes)/2; i++ {
		buffer := stringByRunes[i]
		stringByRunes[i] = stringByRunes[len(stringByRunes)-1-i]
		stringByRunes[len(stringByRunes)-1-i] = buffer
	}
	return string(stringByRunes)
}
