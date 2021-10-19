package finder

import (
	"context"
)

type Finder struct {
}

func NewFinder() *Finder {
	return &Finder{}
}

func (f *Finder) FindDuplicates(_ context.Context, list []string) ([]string, error) {
	listMap := make(map[string]struct{})
	duplicatesMap := make(map[string]struct{})
	for _, s := range list {
		_, ok := listMap[s]
		if ok {
			duplicatesMap[s] = struct{}{}
			continue
		}

		listMap[s] = struct{}{}
	}

	result := make([]string, 0, len(duplicatesMap))
	for key := range duplicatesMap {
		result = append(result, key)
	}

	return result, nil
}
