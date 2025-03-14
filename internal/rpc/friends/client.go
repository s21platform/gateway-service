package friends

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	friends_proto "github.com/s21platform/friends-proto/friends-proto"

	"github.com/s21platform/gateway-service/internal/config"
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
		return nil, fmt.Errorf("failed to s.client.GetCountFriends: %v", err)
	}
	return resp, nil
}

func (s *Service) SetFriends(ctx context.Context, peer *friends_proto.SetFriendsIn) (*friends_proto.SetFriendsOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.client.SetFriends(ctx, peer)
	if err != nil {
		return nil, fmt.Errorf("failed to s.client.SetFriends: %v", err)
	}
	return resp, nil
}

func (s *Service) RemoveFriends(ctx context.Context, peer *friends_proto.RemoveFriendsIn) (*friends_proto.RemoveFriendsOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.client.RemoveFriends(ctx, peer)
	if err != nil {
		return nil, fmt.Errorf("failed to s.client.RemoveFriends: %v", err)
	}
	return resp, nil
}

func (s *Service) CheckSubscribeToPeer(ctx context.Context, peer *friends_proto.IsFriendExistIn) (*friends_proto.IsFriendExistOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.client.IsFriendExist(ctx, peer)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
