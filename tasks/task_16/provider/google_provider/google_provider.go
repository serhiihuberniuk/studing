package google_provider

import (
	"context"
	"fmt"
)

type GoogleProvider struct {
	GoogleDriveProvider  *GoogleDriveProvider
	GoogleSheetsProvider *GoogleSheetsProvider
}

func New(ctx context.Context, pathToConfig string) (*GoogleProvider, error) {
	c, err := ReadConfig(pathToConfig)
	if err != nil {
		return nil, fmt.Errorf("error while reading configs: %w", err)
	}

	drive, err := NewGoogleDriveProvider(ctx, c.CredentialsFile)
	if err != nil {
		return nil, fmt.Errorf("error while creating drive provider: %w", err)
	}

	sheet, err := NewGoogleSheetsProvider(ctx, c.CredentialsFile)
	if err != nil {
		return nil, fmt.Errorf("error while creating sheet proider: %w", err)
	}

	return &GoogleProvider{
		GoogleSheetsProvider: sheet,
		GoogleDriveProvider:  drive,
	}, nil
}
