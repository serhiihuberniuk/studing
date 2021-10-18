package scanner

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

type TerminalScanner struct {
}

func NewTerminalScanner() *TerminalScanner {
	return &TerminalScanner{}
}

func (s *TerminalScanner) Scan(_ context.Context) (string, error) {

	fmt.Print("->")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return strings.Trim(scanner.Text(), "\n"), nil
}
