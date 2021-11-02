package printer

import (
	"context"
	"fmt"

	"studing/tasks/task_16/models"
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
	Parse(ctx context.Context, excelFile *models.ExcelFile) (*models.ExcelFile, error)
}

func (p *ExcelPrinter) Print(ctx context.Context, excelFile *models.ExcelFile) error {

	excelFile, err := p.parser.Parse(ctx, excelFile)
	if err != nil {
		return fmt.Errorf("error while parsing excel file: %w", err)
	}

	for _, sheet := range excelFile.Sheets {
		err := p.PrintSheet(sheet)
		if err != nil {
			return fmt.Errorf("error while printing sheet: %w", err)
		}
	}

	return nil
}

func (p *ExcelPrinter) PrintSheet(sheet *models.Sheet) error {
	var maxCellLen int
	var maxCellsCount int
	for _, row := range sheet.Rows {
		if len(row.Cells) > maxCellsCount {
			maxCellsCount = len(row.Cells)
		}

		for _, cell := range row.Cells {
			if len(cell.Value) > maxCellLen {
				maxCellLen = len(cell.Value)
			}
		}
	}

	fmt.Println(sheet.Name)

	for _, row := range sheet.Rows {
		if len(row.Cells) < maxCellsCount {
			row.Cells = append(row.Cells, make([]*models.Cell, maxCellsCount-len(row.Cells))...)
		}

		printLinesDivider(maxCellsCount*(maxCellLen+1)+1, "=")
		fmt.Println()

		for i, cell := range row.Cells {
			if i == 0 {
				fmt.Print("|")
			}
			if cell == nil {
				cell = &models.Cell{
					Value: "",
				}
			}
			fmt.Print(cell.Value)

			for i := 0; i < (maxCellLen - len([]rune(cell.Value))); i++ {
				fmt.Print(" ")
			}

			fmt.Print("|")
		}

		fmt.Println()
	}

	printLinesDivider(maxCellsCount*(maxCellLen+1)+1, "=")
	fmt.Println()

	return nil
}

func printLinesDivider(len int, divideSymbol string) {
	for i := 0; i < len; i++ {
		fmt.Print(divideSymbol)
	}
}
