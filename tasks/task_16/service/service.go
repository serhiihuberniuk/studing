package service

import (
	"context"
	"fmt"

	"studing/tasks/task_16/models"
)

type Service struct {
	scanner scanner
	parser  parser
	printer printer
}

func New(scanner scanner, parser parser, printer printer) *Service {
	return &Service{
		scanner: scanner,
		printer: printer,
		parser:  parser,
	}
}

type printer interface {
	Print(ctx context.Context, excelFile *models.ExcelFile) error
}

type source interface {
	GetExcelFiles(ctx context.Context) ([]*models.ExcelFile, error)
	Convert(ctx context.Context, excelFile *models.ExcelFile) (func(), error)
}

type scanner interface {
	Scan(ctx context.Context) (string, error)
}

type parser interface {
	Parse(ctx context.Context, excelFile *models.ExcelFile) ([]models.Sheet, error)
}

func (s *Service) OpenExcelFile(ctx context.Context, source source) error {
	files, err := source.GetExcelFiles(ctx)
	if err != nil {
		return fmt.Errorf("error while getting files: %w", err)
	}

	if len(files) == 0 {
		fmt.Println("cannot find any excel files")

		return nil
	}

	file, deleteFunc, err := s.selectFile(ctx, files, source)
	if err != nil {
		return fmt.Errorf("error while selecting file: %w", err)
	}
	defer deleteFunc()

	sheets, err := s.parser.Parse(ctx, file)
	if err != nil {
		return fmt.Errorf("error while parsing file: %w", err)
	}

	file.Sheets = sheets
	if err = s.printer.Print(ctx, file); err != nil {
		return fmt.Errorf("error while printing file: %w", err)
	}

	return nil
}

func (s *Service) selectFile(ctx context.Context, files []*models.ExcelFile, source source) (*models.ExcelFile, func(), error) {
	for _, file := range files {
		fmt.Println(file.Name)
	}

	fmt.Println("Enter name of file you want to read: ")
	fileName, err := s.scanner.Scan(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("error while scanning: %w", err)
	}

	var fileToOpen *models.ExcelFile

	for _, file := range files {
		if file.Name == fileName {
			fileToOpen = file
		}
	}

	if fileToOpen == nil {
		return nil, nil, models.ErrNotFound
	}

	deleteFunc := func() {}
	if fileToOpen.NeedConvert {
		deleteFunc, err = source.Convert(ctx, fileToOpen)
		if err != nil {
			return nil, nil, fmt.Errorf("error while converting file: %w", err)
		}
	}

	return fileToOpen, deleteFunc, nil
}
