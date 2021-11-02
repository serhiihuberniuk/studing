package google_provider

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
	"studing/tasks/task_16/configs"
)

type GoogleProvider struct {
	client *http.Client
}

func NewGoogleProvider(ctx context.Context, c *configs.Config) (*GoogleProvider, error) {
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

	return &GoogleProvider{
		client: client,
	}, nil
}
