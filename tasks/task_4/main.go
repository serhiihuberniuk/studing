package main

import (
	"fmt"
	"strings"
)

func main() {
	text := " a b c d e f g "
	fmt.Println(textFromTheEnd(text))

}

func textFromTheEnd(text string) string {
	var stringBuffer strings.Builder
	stringByRunes := []rune(text)
	var newString []rune
	for _, letter := range stringByRunes {
		newString = append(newString, letter)
	}
	for i := len(newString) - 1; i >= 0; i-- {
		stringBuffer.WriteRune(newString[i])
	}
	return stringBuffer.String()
}
