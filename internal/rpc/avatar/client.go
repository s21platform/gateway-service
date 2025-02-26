package avatar

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	avatar "github.com/s21platform/avatar-proto/avatar-proto"

	"github.com/s21platform/gateway-service/internal/config"
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

func (s *Service) SetUserAvatar(ctx context.Context, filename string, file multipart.File, uuid string) (*avatar.SetUserAvatarOut, error) {
	stream, err := s.client.SetUserAvatar(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to set avatar: %w", err)
	}
	req := avatar.SetUserAvatarIn{
		Filename: filename,
		Uuid:     uuid,
	}
	if err = stream.Send(&req); err != nil {
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
		req = avatar.SetUserAvatarIn{Batch: buf[:n]}
		if err = stream.Send(&req); err != nil {
			return nil, fmt.Errorf("failed to send set avatar: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return nil, fmt.Errorf("failed to receive set avatar response: %v", err)
	}
	return resp, nil
}

func (s *Service) GetAllUserAvatars(ctx context.Context, uuid string) (*avatar.GetAllUserAvatarsOut, error) {
	req := avatar.GetAllUserAvatarsIn{
		Uuid: uuid,
	}
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.client.GetAllUserAvatars(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to get all user avatars: %w", err)
	}

	return resp, nil
}

func (s *Service) DeleteUserAvatar(ctx context.Context, id int32) (*avatar.Avatar, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))

	req := avatar.DeleteUserAvatarIn{
		AvatarId: id,
	}

	resp, err := s.client.DeleteUserAvatar(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to delete user avatar: %w", err)
	}

	return resp, nil
}

func (s *Service) SetSocietyAvatar(ctx context.Context, filename string, file multipart.File, uuid string) (*avatar.SetSocietyAvatarOut, error) {
	stream, err := s.client.SetSocietyAvatar(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to set avatar: %w", err)
	}
	req := avatar.SetSocietyAvatarIn{
		Filename: filename,
		Uuid:     uuid,
	}
	if err = stream.Send(&req); err != nil {
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
		req = avatar.SetSocietyAvatarIn{Batch: buf[:n]}
		if err = stream.Send(&req); err != nil {
			return nil, fmt.Errorf("failed to send set avatar: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return nil, fmt.Errorf("failed to receive set avatar response: %v", err)
	}
	return resp, nil
}

func (s *Service) GetAllSocietyAvatars(ctx context.Context, uuid string) (*avatar.GetAllSocietyAvatarsOut, error) {
	req := avatar.GetAllSocietyAvatarsIn{
		Uuid: uuid,
	}
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.client.GetAllSocietyAvatars(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to get all society avatars: %w", err)
	}

	return resp, nil
}

func (s *Service) DeleteSocietyAvatar(ctx context.Context, id int32) (*avatar.Avatar, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))

	req := avatar.DeleteSocietyAvatarIn{
		AvatarId: id,
	}

	resp, err := s.client.DeleteSocietyAvatar(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to delete society avatar: %w", err)
	}

	return resp, nil
}
