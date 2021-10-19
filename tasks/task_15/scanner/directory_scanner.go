package scanner

import (
	"context"
	"fmt"
	"os"
	"path"
)

type DirectoryScanner struct {
}

func NewDirectoryScanner() *DirectoryScanner {
	return &DirectoryScanner{}
}

func (ds *DirectoryScanner) ScanRecursively(ctx context.Context, pathToDirectory string) ([]string, error) {
	files, err := os.ReadDir(pathToDirectory)
	if err != nil {
		return nil, fmt.Errorf("error while reading directory: %w", err)
	}

	var fileNames []string
	for _, file := range files {
		if file.IsDir() {
			fileNamesInChild, err := ds.ScanRecursively(ctx, path.Join(pathToDirectory, file.Name()))
			if err != nil {
				return nil, fmt.Errorf("error while scanning directory: %w", err)
			}

			fileNames = append(fileNames, fileNamesInChild...)

			continue
		}

		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}
