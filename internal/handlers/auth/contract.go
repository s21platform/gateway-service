package auth

import (
	"context"
	"github.com/s21platform/gateway-service/internal/rpc/auth"
)

type Usecase interface {
	Login(ctx context.Context, username string, password string) (*auth.JWT, error)
}
