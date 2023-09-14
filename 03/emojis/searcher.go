package main

import (
	"context"
	"errors"
	"slices"
	"sort"
	"strings"

	"github.com/ServiceWeaver/weaver"
)

type Searcher interface {
	Search(ctx context.Context, query string) ([]string, error)
}

type searcher struct {
	weaver.Implements[Searcher]
}

func (s *searcher) Search(_ context.Context, query string) ([]string, error) {
	lcQueryWords := strings.Fields(strings.ToLower(query))

	if len(lcQueryWords) == 0 {
		return []string{}, errors.New("Query must contain at least one word")
	}

	results := make([]string, 0)

	for emoji, labels := range emojis {
		foundAllWords := true

		for _, word := range lcQueryWords {
			if !slices.Contains(labels, word) {
				foundAllWords = false
				break
			}
		}

		if foundAllWords {
			results = append(results, emoji)
		}
	}

	sort.Strings(results)

	return results, nil
}
