package user

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	user "github.com/s21platform/user-service/pkg/user"

	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/model"
)

type Service struct {
	clientUser user.UserServiceClient
}

func NewService(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.User.Host, cfg.User.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	client := user.NewUserServiceClient(conn)
	return &Service{clientUser: client}
}

func (s *Service) GetInfo(ctx context.Context, uuid string) (*user.GetUserInfoByUUIDOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.clientUser.GetUserInfoByUUID(ctx, &user.GetUserInfoByUUIDIn{Uuid: uuid})
	if err != nil {
		log.Printf("failed to call: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Service) UpdateProfile(ctx context.Context, data model.ProfileData) (*user.UpdateProfileOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.clientUser.UpdateProfile(ctx, data.FromDTO())
	if err != nil {
		return nil, fmt.Errorf("failed to update user profile: %v", err)
	}
	return resp, nil
}

func (s *Service) GetCountFriends(ctx context.Context) (*user.GetCountFriendsOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.clientUser.GetCountFriends(ctx, &user.EmptyFriends{})
	if err != nil {
		return nil, fmt.Errorf("failed to s.client.GetCountFriends: %v", err)
	}
	return resp, nil
}

func (s *Service) SetFriends(ctx context.Context, peer *user.SetFriendsIn) (*user.SetFriendsOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.clientUser.SetFriends(ctx, peer)
	if err != nil {
		return nil, fmt.Errorf("failed to s.client.SetFriends: %v", err)
	}
	return resp, nil
}

func (s *Service) RemoveFriends(ctx context.Context, peer *user.RemoveFriendsIn) (*user.RemoveFriendsOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.clientUser.RemoveFriends(ctx, peer)
	if err != nil {
		return nil, fmt.Errorf("failed to s.client.RemoveFriends: %v", err)
	}
	return resp, nil
}
