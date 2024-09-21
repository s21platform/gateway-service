package auth

import (
	"context"
	"github.com/s21platform/gateway-service/internal/rpc/auth"
)

type UseCase struct {
	aC AuthClient
}

func New(aC AuthClient) *UseCase {
	return &UseCase{aC: aC}
}

func (uc *UseCase) Login(ctx context.Context, username string, password string) (*auth.JWT, error) {
	return uc.aC.DoLogin(ctx, username, password)
}
