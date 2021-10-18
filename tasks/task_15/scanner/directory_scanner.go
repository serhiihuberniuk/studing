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

func (ds *DirectoryScanner) ScanRecursively(ctx context.Context, path string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("error while reading directory: %w", err)
	}

	var allFiles []os.DirEntry
	for _, file := range files {
		if file.IsDir() {
			path = path + "/"
			fileNamesInChild, err := ds.ScanRecursively(ctx, path+file.Name())
			if err != nil {
				return nil, fmt.Errorf("error while scanning directory: %w", err)
			}

			allFiles = append(allFiles, fileNamesInChild...)

			continue
		}

		allFiles = append(allFiles, file)
	}

	return allFiles, nil
}
