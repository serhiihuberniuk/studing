package printer

import (
	"context"
	"fmt"
)

type ExcelPrinter struct {
	parser parser
}

func NewExcelPrinter(p parser) *ExcelPrinter {
	return &ExcelPrinter{
		parser: p,
	}
}

type parser interface {
	Parse(ctx context.Context, fileName string) (map[string][][]string, error)
}

func (p *ExcelPrinter) PrintExcelFile(ctx context.Context, fileName string) error {
	excelMap, err := p.parser.Parse(ctx, fileName)
	if err != nil {
		return fmt.Errorf("error while parsing excel: %w", err)
	}

	for sheet, rows := range excelMap {
		err := p.printSheet(ctx, sheet, rows)
		if err != nil {
			return fmt.Errorf("error while printing file: %w", err)
		}
	}

	return nil
}

func (p *ExcelPrinter) printSheet(_ context.Context, name string, rows [][]string) error {
	var maxCellLen int
	var maxCellCount int
	for _, row := range rows {
		if len(row) > maxCellCount {
			maxCellCount = len(row)
		}
		for _, cell := range row {
			if len(cell) > maxCellLen {
				maxCellLen = len(cell)
			}
		}
	}

	fmt.Println(name)
	for _, row := range rows {
		if len(row) < maxCellCount {
			row = append(row, make([]string, maxCellCount-len(row))...)
		}

		printLinesDivider(maxCellCount*(maxCellLen+1)+1, "=")
		fmt.Println()

		for i, cell := range row {
			if i == 0 {
				fmt.Print("|")
			}

			fmt.Print(cell)

			for i := 0; i < (maxCellLen - len([]rune(cell))); i++ {
				fmt.Print(" ")
			}

			fmt.Print("|")
		}

		fmt.Println()
	}

	printLinesDivider(maxCellCount*(maxCellLen+1)+1, "=")
	fmt.Println()

	return nil
}

func printLinesDivider(len int, divideSymbol string) {
	for i := 0; i < len; i++ {
		fmt.Print(divideSymbol)
	}
}
