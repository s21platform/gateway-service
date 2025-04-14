package feed

import (
	"context"

	feed "github.com/s21platform/feed-proto/feed-proto"
)

type FeedClient interface {
	CreateUserPost(ctx context.Context, content string) (*feed.CreateUserPostOut, error)
}
