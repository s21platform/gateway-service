package search

import (
	"context"
	"github.com/s21platform/search-proto/search"
)

type SearchClient interface {
	SearchSociety(ctx context.Context, searchName string) (*search.GetSocietyOut, error)
}
