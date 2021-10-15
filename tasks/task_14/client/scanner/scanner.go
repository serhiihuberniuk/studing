package scanner

import (
	"bufio"
	"fmt"
	"os"
)

type Scanner struct {
}

func NewScanner() *Scanner {
	return &Scanner{}
}

func (s *Scanner) ScanTerminal() (string, error) {

	fmt.Print("->")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text(), nil
}
