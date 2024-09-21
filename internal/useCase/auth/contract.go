package auth

import (
	"context"
	"github.com/s21platform/gateway-service/internal/rpc/auth"
)

type AuthClient interface {
	DoLogin(ctx context.Context, username, password string) (*auth.JWT, error)
}
