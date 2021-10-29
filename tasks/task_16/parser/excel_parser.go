package parser

import (
	"context"
	"fmt"

	"github.com/xuri/excelize/v2"
)

type ExcelParser struct {
}

func NewExcelParser() *ExcelParser {
	return &ExcelParser{}
}

func (p *ExcelParser) Parse(_ context.Context, filename string) (map[string][][]string, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error occured while opening excel file: %w", err)
	}

	sheets := f.GetSheetList()
	sheetsMap := make(map[string][][]string)
	for _, sheet := range sheets {
		rows, err := f.GetRows(sheet)
		if err != nil {
			return nil, fmt.Errorf("error while getting rows from sheet: %w", err)
		}

		sheetsMap[sheet] = rows
	}

	return sheetsMap, nil
}
