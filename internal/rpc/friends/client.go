package friends

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/metadata"

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

func (s *Service) GetCountFriends(ctx context.Context) (*friends_proto.GetCountFriendsOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.client.GetCountFriends(ctx, &friends_proto.EmptyFriends{})
	if err != nil {
		return nil, fmt.Errorf("s.client.GetCountFriends: %v", err)
	}
	log.Println("resp: ", resp)
	return resp, nil
}
