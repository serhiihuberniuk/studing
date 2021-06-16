package main

import (
	"github.com/rs/zerolog/log"
)

func message(s string) {
	log.Print(s)
}
func main() {
	s := "message"
	message(s)

}
