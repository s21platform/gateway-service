package auth

import (
	"context"

	authproto "github.com/s21platform/auth-service/pkg/auth"

	"github.com/s21platform/gateway-service/internal/rpc/auth"
)

type AuthClient interface {
	DoLogin(ctx context.Context, username, password string) (*auth.JWT, error)
	CheckEmailAvailability(ctx context.Context, email string) (*authproto.CheckEmailAvailabilityOut, error)
}
