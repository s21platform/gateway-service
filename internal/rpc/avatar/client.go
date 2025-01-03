package avatar

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	avatar "github.com/s21platform/avatar-proto/avatar-proto"
	"github.com/s21platform/gateway-service/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type Service struct {
	client avatar.AvatarServiceClient
}

func New(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Avatar.Host, cfg.Avatar.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create grpc client: %v", err)
	}
	client := avatar.NewAvatarServiceClient(conn)
	return &Service{client: client}
}

func (s *Service) SetAvatar(ctx context.Context, filename string, file multipart.File, uuid string) (*avatar.SetAvatarOut, error) {
	stream, err := s.client.SetAvatar(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to set avatar: %w", err)
	}
	req := avatar.SetAvatarIn{
		Filename: filename,
		UserUuid: uuid,
	}
	if err := stream.Send(&req); err != nil {
		return nil, fmt.Errorf("failed to send set avatar: %v", err)
	}
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("failed to read file: %v", err)
		}
		req := avatar.SetAvatarIn{Batch: buf[:n]}
		if err := stream.Send(&req); err != nil {
			return nil, fmt.Errorf("failed to send set avatar: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return nil, fmt.Errorf("failed to receive set avatar response: %v", err)
	}
	return resp, nil
}

func (s *Service) GetAllAvatars(ctx context.Context, uuid string) (*avatar.GetAllAvatarsOut, error) {
	req := avatar.GetAllAvatarsIn{
		UserUuid: uuid,
	}
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.client.GetAllAvatars(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to get all avatars: %w", err)
	}

	return resp, nil
}

func (s *Service) DeleteAvatar(ctx context.Context, id int32) (*avatar.Avatar, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))

	req := avatar.DeleteAvatarIn{
		AvatarId: id,
	}

	resp, err := s.client.DeleteAvatar(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to delete avatar: %w", err)
	}

	return resp, nil
}
