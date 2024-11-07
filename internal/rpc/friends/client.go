package friends

import (
	"context"
	"fmt"
	"log"

	friends_proto "github.com/s21platform/friends-proto/friends-proto"
	"github.com/s21platform/gateway-service/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	client friends_proto.FriendsServiceClient
}

func NewService(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Friends.Host, cfg.Friends.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create gRPC client: %v", err)
	}
	client := friends_proto.NewFriendsServiceClient(conn)
	return &Service{client: client}
}

func (s *Service) GetCountFriends(ctx context.Context, req *friends_proto.Empty) (*friends_proto.GetCountFriendsOut, error) {
	resp, err := s.client.GetCountFriends(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("s.client.GetCountFriends: %v", err)
	}
	return resp, nil
}