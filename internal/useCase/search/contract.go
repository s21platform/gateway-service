package search

import (
	"context"

	"github.com/s21platform/search-proto/search"
)

type SearchClient interface {
	GetUserWithLimit(ctx context.Context, in *search.GetUserWithLimitIn) (*search.GetUserWithLimitOut, error)
}