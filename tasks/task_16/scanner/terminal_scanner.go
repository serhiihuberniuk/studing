package scanner

import (
	"bufio"
	"fmt"
	"os"
)

func ScanTerminal() string {
	fmt.Print("->")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}
