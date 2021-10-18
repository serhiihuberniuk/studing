package checker

import (
	"context"
	"fmt"
	"os"
)

type DirectoryChecker struct {
	scanner scanner
	finder  finder
}

type scanner interface {
	ScanRecursively(ctx context.Context, path string) ([]os.DirEntry, error)
}

type finder interface {
	FindDuplicates(ctx context.Context, list []string) ([]string, error)
}

func NewDirectoryChecker(s scanner, f finder) *DirectoryChecker {
	return &DirectoryChecker{
		scanner: s,
		finder:  f,
	}
}

func (dc *DirectoryChecker) CheckDirectoryForDuplicates(ctx context.Context, path string) ([]string, error) {
	files, err := dc.scanner.ScanRecursively(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("error while scanning directory: %w", err)
	}
	if len(files) == 0 {
		fmt.Println("directory is empty")

		return nil, nil
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	duplicates, err := dc.finder.FindDuplicates(ctx, fileNames)
	if err != nil {
		return nil, fmt.Errorf("error while finding duplicates: %w", err)
	}

	return duplicates, nil
}
