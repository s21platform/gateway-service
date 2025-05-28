package feed

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	feedproto "github.com/s21platform/feed-proto/feed-proto"

	"github.com/s21platform/gateway-service/internal/config"
)

type Service struct {
	client feedproto.FeedServiceClient
}

func New(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Feed.Host, cfg.Feed.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create grpc client: %v", err)
	}
	client := feedproto.NewFeedServiceClient(conn)
	return &Service{client: client}
}

// func (s *Service) CreateUserPost(ctx context.Context, content string) (*feedproto.CreateUserPostOut, error) {
// 	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
// 	request := &feedproto.CreateUserPostIn{
// 		Content: content,
// 	}
// 	resp, err := s.client.CreateUserPost(ctx, request)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create post in grpc: %w", err)
// 	}
// 	return resp, nil
// }
