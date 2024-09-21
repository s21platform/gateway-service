package auth

import (
	"context"
	"fmt"
	auth "github.com/s21platform/auth-proto/auth-proto"
	"github.com/s21platform/gateway-service/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Service struct {
	client auth.AuthServiceClient
}

type JWT struct {
	Jwt string `json:"jwt"`
}

func NewService(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Auth.Host, cfg.Auth.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to create grpc connection: %v", err)
	}
	client := auth.NewAuthServiceClient(conn)
	return &Service{client: client}
}

func (s *Service) DoLogin(ctx context.Context, username, password string) (*JWT, error) {
	resp, err := s.client.Login(ctx, &auth.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		if statusError, ok := status.FromError(err); ok {
			switch statusError.Code() {
			case codes.InvalidArgument:
				return nil, status.Error(codes.InvalidArgument, "Неверно введены логин или пароль")
			default:
				return nil, status.Error(codes.Internal, "Неизвестная ошибка")
			}
		}
		return nil, status.Error(codes.Internal, "Unknown error")
	}
	return &JWT{Jwt: resp.Jwt}, nil
}
