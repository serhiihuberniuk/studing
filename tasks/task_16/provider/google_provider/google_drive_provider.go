package google_provider

import (
	"context"
	"fmt"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"studing/tasks/task_16/models"
)

const (
	googleSheetMimeType = "application/vnd.google-apps.spreadsheet"
	excelMimeType       = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
)

type GoogleDriveProvider struct {
	srv *drive.Service
}

func NewGoogleDriveProvider(ctx context.Context, credentialsFile string) (*GoogleDriveProvider, error) {
	srv, err := drive.NewService(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, fmt.Errorf("error while creating server: %w", err)
	}

	return &GoogleDriveProvider{
		srv: srv,
	}, nil
}

func (p *GoogleDriveProvider) getFilesList(ctx context.Context) ([]*drive.File, error) {

	fileList, err := p.srv.Files.List().Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("error while getting list from google drive : %w", err)
	}

	return fileList.Files, nil
}

func (p *GoogleDriveProvider) GetExcelFiles(ctx context.Context) ([]*models.ExcelFile, error) {
	list, err := p.getFilesList(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while getting list of files from google drive: %w", err)
	}

	var excelFiles []*models.ExcelFile
	for _, file := range list {
		switch file.MimeType {
		case googleSheetMimeType:
			excelFiles = append(excelFiles, &models.ExcelFile{
				Name: file.Name,
				ID:   file.Id,
			})
		case excelMimeType:
			excelFiles = append(excelFiles, &models.ExcelFile{
				Name:        file.Name,
				ID:          file.Id,
				NeedConvert: true,
			})
		}
	}

	return excelFiles, nil
}

func (p *GoogleDriveProvider) Convert(ctx context.Context, file *models.ExcelFile) error {
	w, err := p.srv.Files.Get(file.ID).Context(ctx).Download()
	if err != nil {
		return fmt.Errorf("error while getting file: %w", err)
	}
	defer w.Body.Close()

	fileMetadata := &drive.File{
		Name:     file.Name,
		MimeType: googleSheetMimeType,
	}

	f, err := p.srv.Files.Create(fileMetadata).Media(w.Body).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("error while creating file: %w", err)
	}

	file.ID = f.Id

	return nil

}

func (p *GoogleDriveProvider) Delete(ctx context.Context, file *models.ExcelFile) error {
	if err := p.srv.Files.Delete(file.ID).Context(ctx).Do(); err != nil {
		return fmt.Errorf("error while deleting file: %w", err)
	}

	return nil
}
