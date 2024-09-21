package auth

import (
	"context"
	"github.com/s21platform/gateway-service/internal/rpc/auth"
)

type AuthUsecase interface {
	Login(ctx context.Context, username string, password string) (*auth.JWT, error)
}
