package feed

import (
	"context"
	"fmt"
	"log"

	feed "github.com/s21platform/feed-proto/feed-proto"
	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	client feed.FeedServiceClient
}

func New(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Feed.Host, cfg.Feed.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create grpc client: %v", err)
	}
	client := feed.NewFeedServiceClient(conn)
	return &Service{client: client}
}

func (s *Service) CreateUserPost(ctx context.Context, req *model.CreateUserPostRequestData) (*feed.CreateUserPostOut, error) {

}
