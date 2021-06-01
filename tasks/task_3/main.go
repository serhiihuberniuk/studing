package main

import "fmt"

func deleteGaps(text string) {
	var previousLetter bool
	for i := 0; i < len(text); i++ {
		currentLetter := text[i] == ' '
		if currentLetter {
			if !previousLetter {
				fmt.Printf("%c", text[i])
			}
		} else {
			fmt.Printf("%c", text[i])
		}
		previousLetter = currentLetter
	}
}
func main() {
	text := "   a b  c                     d e  "
	deleteGaps(text)

}
