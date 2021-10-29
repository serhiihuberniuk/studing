package provider

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/option"
	"studing/tasks/task_16/configs"
)

type GoogleDriveProvider struct {
	srv *drive.Service
}

func NewGoogleDriveProvider(ctx context.Context, c *configs.Config) (*GoogleDriveProvider, error) {
	b, err := ioutil.ReadFile(c.CredentialsFile)
	if err != nil {
		return nil, fmt.Errorf("error while reading credential JSON file: %w", err)
	}

	authConfig, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		return nil, fmt.Errorf("error while constructing configs: %w", err)
	}

	client, err := getClient(ctx, authConfig, c.TokenFile)
	if err != nil {
		return nil, fmt.Errorf("error while getting client: %w", err)
	}

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("error while creating server: %w", err)
	}

	return &GoogleDriveProvider{
		srv: srv,
	}, nil
}

func (p *GoogleDriveProvider) getFilesList() ([]*drive.File, error) {

	fileList, err := p.srv.Files.List().Do()
	if err != nil {
		return nil, fmt.Errorf("error while getting list from google drive : %w", err)
	}

	files := fileList.Items

	return files, nil
}

func (p *GoogleDriveProvider) GetExcelFilesNames(_ context.Context) ([]string, error) {
	list, err := p.getFilesList()
	if err != nil {
		return nil, fmt.Errorf("error while getting list of files from google drive: %w", err)
	}

	var excelFiles []string
	for _, file := range list {
		if file.FileExtension == "xlsx" {
			excelFiles = append(excelFiles, file.Title)
		}
	}

	return excelFiles, nil
}

func (p *GoogleDriveProvider) DownloadFile(_ context.Context, name string, downloadTo *os.File) error {
	list, err := p.getFilesList()
	if err != nil {
		return fmt.Errorf("error while getting list of files from google drive: %w", err)
	}

	var fileToDownload *drive.File
	for _, file := range list {
		if file.Title == name {
			fileToDownload = file
		}
	}

	if fileToDownload == nil {
		return errors.New("cannot find file with such name")
	}

	w, err := p.srv.Files.Get(fileToDownload.Id).Download()
	if err != nil {
		return fmt.Errorf("error while downloading file: %w", err)
	}
	defer w.Body.Close()

	_, err = bufio.NewReader(w.Body).WriteTo(downloadTo)
	if err != nil {
		return fmt.Errorf("error while writing to destination file: %w", err)
	}

	return nil
}
