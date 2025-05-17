package auth

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/s21platform/auth-service/pkg/auth"

	"github.com/s21platform/gateway-service/internal/config"
)

type Service struct {
	client auth.AuthServiceClient
}

type JWT struct {
	Jwt string `json:"jwt"`
}

func NewService(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Auth.Host, cfg.Auth.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

func (s *Service) CheckEmailAvailability(ctx context.Context, email string) (*auth.CheckEmailAvailabilityOut, error) {
	resp, err := s.client.CheckEmailAvailability(ctx, &auth.CheckEmailAvailabilityIn{Email: email})
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to check email in grpc: %v", err))
	}

	return resp, nil
}

func (s *Service) SendUserVerificationCode(ctx context.Context, email string) (*auth.SendUserVerificationCodeOut, error) {
	resp, err := s.client.SendUserVerificationCode(ctx, &auth.SendUserVerificationCodeIn{Email: email})
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to send code in grpc: %v", err))
	}

	return resp, nil
}

func (s *Service) LoginV2(ctx context.Context, login, password string) (*auth.LoginV2Out, error) {
	resp, err := s.client.LoginV2(ctx, &auth.LoginV2In{
		Login:    login,
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
	return resp, nil
}
