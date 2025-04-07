package feed

import (
	"context"

	feed "github.com/s21platform/feed-proto/feed-proto"
	"github.com/s21platform/gateway-service/internal/model"
)

type FeedClient interface {
	CreateUserPost(ctx context.Context, req *model.CreateUserPostRequestData) (*feed.CreateUserPostOut, error)
}
