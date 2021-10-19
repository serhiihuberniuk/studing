package scanner

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

type TerminalScanner struct {
}

func NewTerminalScanner() *TerminalScanner {
	return &TerminalScanner{}
}

func (ts *TerminalScanner) Scan(_ context.Context) (string, error) {

	fmt.Print("->")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text(), nil

}
