package google_provider

import (
	"context"
	"errors"
	"fmt"
	"log"

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

func NewGoogleDriveProvider(ctx context.Context, gp *GoogleProvider) (*GoogleDriveProvider, error) {
	srv, err := drive.NewService(ctx, option.WithHTTPClient(gp.client))
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
	files := fileList.Files

	return files, nil
}

func (p *GoogleDriveProvider) GetExcelFiles(ctx context.Context) ([]*models.ExcelFile, error) {
	list, err := p.getFilesList(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while getting list of files from google drive: %w", err)
	}

	var excelFiles []*models.ExcelFile
	for _, file := range list {
		if file.MimeType == excelMimeType || file.MimeType == googleSheetMimeType {
			excelFiles = append(excelFiles, &models.ExcelFile{
				Name: file.Name,
				ID:   file.Id,
			})
		}
	}

	if len(excelFiles) == 0 {
		return nil, errors.New("cannot find any excel files")
	}

	return excelFiles, nil
}

func (p *GoogleDriveProvider) GetFileByName(ctx context.Context, name string) (*models.ExcelFile, error) {
	list, err := p.getFilesList(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while getting list of files from google drive: %w", err)
	}

	for _, file := range list {
		if file.Name == name {
			if file.MimeType == googleSheetMimeType {
				return &models.ExcelFile{Name: file.Name, ID: file.Id}, nil
			}

			f, err := p.ConvertExcelToGoogleSheet(ctx, &models.ExcelFile{Name: file.Name, ID: file.Id})
			if err != nil {
				return nil, fmt.Errorf("error while converting file: %w", err)
			}

			return f, nil
		}
	}

	return nil, errors.New("cannot file with such name")
}

func (p *GoogleDriveProvider) ConvertExcelToGoogleSheet(ctx context.Context, file *models.ExcelFile) (*models.ExcelFile, error) {
	w, err := p.srv.Files.Get(file.ID).Context(ctx).Download()
	if err != nil {
		return nil, fmt.Errorf("error while downloading file: %w", err)
	}
	defer w.Body.Close()

	fileMetadata := &drive.File{
		Name:     file.Name,
		MimeType: googleSheetMimeType,
	}

	f, err := p.srv.Files.Create(fileMetadata).Media(w.Body).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("error while creating file: %w", err)
	}
	err = p.srv.Files.Delete(file.ID).Context(ctx).Do()
	if err != nil {
		log.Printf(fmt.Errorf("converted file has not been deleted: %w", err).Error())
	}

	return &models.ExcelFile{
		Name: f.Name,
		ID:   f.Id,
	}, nil

}
