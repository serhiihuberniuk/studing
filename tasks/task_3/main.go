package main

import "fmt"

func task3(text string) (newText string) {
	if text[0] != ' ' {
		fmt.Printf("%c", text[0])
	}
	for i := 1; i < len(text); i++ {
		if (text[i] == ' ' && text[i-1] != ' ') || (text[i] != ' ') {

		}
	}

}
func main() {
	text := "   a b  c                     d e  "
	task3(text)

}
