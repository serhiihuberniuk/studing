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
	listMap := make(map[string]int)
	for _, s := range list {
		listMap[s]++
	}

	var result []string
	for key, v := range listMap {
		if v > 1 {
			result = append(result, key)
		}
	}

	return result, nil
}
