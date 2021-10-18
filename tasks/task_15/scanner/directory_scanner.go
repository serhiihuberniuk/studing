package scanner

import (
	"context"
	"fmt"
	"os"
)

type DirectoryScanner struct {
}

func NewDirectoryScanner() *DirectoryScanner {
	return &DirectoryScanner{}
}

func (ds *DirectoryScanner) ScanRecursively(ctx context.Context, path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("error while reading directory: %w", err)
	}

	var fileNames []string
	for _, file := range files {
		if file.IsDir() {
			path = path + string(os.PathSeparator)
			fileNamesInChild, err := ds.ScanRecursively(ctx, path+file.Name())
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
