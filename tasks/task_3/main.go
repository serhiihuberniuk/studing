package main

import (
	"bytes"
	"fmt"
)

func deleteGaps(text string) string {
	var previousLetter bool
	var byteBuffer bytes.Buffer
	for _, letter := range text {
		currentLetter := letter == ' '
		if currentLetter {
			if !previousLetter {
				byteBuffer.WriteRune(letter)
			}
		} else {
			byteBuffer.WriteRune(letter)
		}
		previousLetter = currentLetter
	}
	return byteBuffer.String()

}
func main() {
	text := "a b  c                     d e  "
	fmt.Println(deleteGaps(text))

}
