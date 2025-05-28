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

	"github.com/s21platform/user-service/pkg/user"

	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/model"
)

type Service struct {
	client user.UserServiceClient
}

func NewService(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.User.Host, cfg.User.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	client := user.NewUserServiceClient(conn)
	return &Service{client: client}
}

func (s *Service) GetInfo(ctx context.Context, uuid string) (*user.GetUserInfoByUUIDOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.client.GetUserInfoByUUID(ctx, &user.GetUserInfoByUUIDIn{
		Uuid: uuid,
	})
	if err != nil {
		log.Printf("failed to call: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Service) UpdateProfile(ctx context.Context, data model.ProfileData) (*user.UpdateProfileOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.client.UpdateProfile(ctx, data.FromDTO())
	if err != nil {
		return nil, fmt.Errorf("failed to update user profile: %v", err)
	}
	return resp, nil
}

func (s *Service) CreatePost(ctx context.Context, content string) (*user.CreatePostOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	request := &user.CreatePostIn{
		Content: content,
	}
	resp, err := s.client.CreatePost(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to create post in grpc: %w", err)
	}
	return resp, nil
}