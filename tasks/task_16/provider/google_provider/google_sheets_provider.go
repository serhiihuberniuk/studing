package google_provider

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"studing/tasks/task_16/models"
)

type GoogleSheetsProvider struct {
	srv *sheets.Service
}

func NewGoogleSheetsProvider(ctx context.Context, credentialsFile string) (*GoogleSheetsProvider, error) {
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, fmt.Errorf("error while crating service: %w", err)
	}

	return &GoogleSheetsProvider{
		srv: srv,
	}, nil
}

func (p *GoogleSheetsProvider) Parse(ctx context.Context, excelFile *models.ExcelFile) error {
	s, err := p.srv.Spreadsheets.Get(excelFile.ID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("error while getting sheets: %w", err)
	}

	for _, v := range s.Sheets {
		sheet := models.Sheet{Name: v.Properties.Title}

		vr, err := p.srv.Spreadsheets.Values.Get(excelFile.ID, sheet.Name).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("error while getting values from sheet: %w", err)
		}

		p.parseSheet(vr.Values, &sheet)

		excelFile.Sheets = append(excelFile.Sheets, sheet)
	}

	return nil
}

func (p *GoogleSheetsProvider) parseSheet(rows [][]interface{}, mySheet *models.Sheet) {
	mySheet.Rows = make([]models.Row, 0, len(rows))

	for _, row := range rows {
		myRow := models.Row{}
		myRow.Cells = make([]models.Cell, 0, len(row))

		for _, cell := range row {
			c := models.Cell{
				Value: cell.(string),
			}

			myRow.Cells = append(myRow.Cells, c)
		}

		mySheet.Rows = append(mySheet.Rows, myRow)
	}
}
