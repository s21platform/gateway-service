package search

import (
	"context"
	"fmt"
	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/search-proto/search"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
)

type Service struct {
	client search.SearchServiceClient
}

func NewService(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Search.Host, cfg.Search.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create search service client: %v", err)
	}
	client := search.NewSearchServiceClient(conn)
	return &Service{client: client}
}

func (s *Service) GetUserWithLimit(ctx context.Context, in *search.GetUserWithLimitIn) (*search.GetUserWithLimitOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.client.GetUserWithLimit(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("failed to get user with limit by search service: %v", err)
	}
	return resp, nil
}
