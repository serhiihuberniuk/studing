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
	listMap := make(map[string]bool)
	var result []string
	for _, s := range list {
		moreThenOnce, ok := listMap[s]
		if ok {
			if moreThenOnce != true {
				result = append(result, s)

			}

			listMap[s] = true
			continue
		}
		listMap[s] = false
	}

	return result, nil
}
