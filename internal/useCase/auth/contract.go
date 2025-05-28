package auth

import (
	"context"

	authproto "github.com/s21platform/auth-service/pkg/auth"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/s21platform/gateway-service/internal/model"
	"github.com/s21platform/gateway-service/internal/rpc/auth"
)

type AuthClient interface {
	DoLogin(ctx context.Context, username, password string) (*auth.JWT, error)
	CheckEmailAvailability(ctx context.Context, email string) (*authproto.CheckEmailAvailabilityOut, error)
	SendUserVerificationCode(ctx context.Context, email string) (*authproto.SendUserVerificationCodeOut, error)
	RegisterUser(ctx context.Context, requestData *model.RegisterRequest) (*emptypb.Empty, error)
	LoginV2(ctx context.Context, login, password string) (*authproto.LoginV2Out, error)
}
