//go:build !test

package auth

import (
	"fmt"
	"github.com/s21platform/gateway-service/internal/config"
	"google.golang.org/grpc"
	"log"
)

type AuthService struct {
	Conn *grpc.ClientConn
}

func NewAuthServiceClient(cfg *config.Config) *AuthService {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.Auth.Host, cfg.Auth.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot dial to auth service: %s", err)
	}
	return &AuthService{Conn: conn}
}
