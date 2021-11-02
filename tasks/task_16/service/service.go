package service

import (
	"context"
	"fmt"

	"studing/tasks/task_16/models"
)

type Service struct {
	scanner scanner
	printer printer
}

func New(s scanner, p printer) *Service {
	return &Service{
		scanner: s,
		printer: p,
	}
}

type printer interface {
	Print(ctx context.Context, excelFile *models.ExcelFile) error
}

type source interface {
	GetExcelFiles(ctx context.Context) ([]*models.ExcelFile, error)
	GetFileByName(ctx context.Context, name string) (*models.ExcelFile, error)
}

type scanner interface {
	Scan(ctx context.Context) (string, error)
}

func (s *Service) OpenExcelFile(ctx context.Context, source source) error {
	files, err := source.GetExcelFiles(ctx)
	if err != nil {
		return fmt.Errorf("error while getting files: %w", err)
	}

	file, err := s.selectFile(ctx, files, source)
	if err != nil {
		return fmt.Errorf("error while selecting file: %w", err)
	}

	if err = s.printer.Print(ctx, file); err != nil {
		return fmt.Errorf("error while printing file: %w", err)
	}

	return nil
}

func (s *Service) selectFile(ctx context.Context, files []*models.ExcelFile, source source) (*models.ExcelFile, error) {
	for _, file := range files {
		fmt.Println(file.Name)
	}

	fmt.Println("Enter name of file you want to read: ")
	fileName, err := s.scanner.Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while scanning: %w", err)
	}

	file, err := source.GetFileByName(ctx, fileName)
	if err != nil {
		return nil, fmt.Errorf("error while getting file by name: %w", err)
	}

	return file, nil
}
