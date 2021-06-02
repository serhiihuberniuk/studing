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
	var stringByRunes []rune
	for _, letter := range text {
		stringByRunes = append(stringByRunes, letter)
	}
	for i := (len(stringByRunes) - 1); i >= 0; i-- {
		stringBuffer.WriteRune(stringByRunes[i])
	}
	return stringBuffer.String()
}
