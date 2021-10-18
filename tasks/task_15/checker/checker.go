package checker

import (
	"context"
	"fmt"
)

type DirectoryChecker struct {
	scanner scanner
	finder  finder
}

type scanner interface {
	ScanRecursively(ctx context.Context, path string) ([]string, error)
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
	fileNames, err := dc.scanner.ScanRecursively(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("error while scanning directory: %w", err)
	}
	if len(fileNames) == 0 {
		fmt.Println("directory is empty")

		return nil, nil
	}

	duplicates, err := dc.finder.FindDuplicates(ctx, fileNames)
	if err != nil {
		return nil, fmt.Errorf("error while finding duplicates: %w", err)
	}

	return duplicates, nil
}
