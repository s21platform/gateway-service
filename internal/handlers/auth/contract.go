package auth

import (
	"context"
	"net/http"

	authproto "github.com/s21platform/auth-service/pkg/auth"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/s21platform/gateway-service/internal/model"
	"github.com/s21platform/gateway-service/internal/rpc/auth"
)

type Usecase interface {
	Login(ctx context.Context, username string, password string) (*auth.JWT, error)
	CheckEmailAvailability(r *http.Request) (*model.EmailResponse, error)
	SendUserVerificationCode(r *http.Request) (*authproto.SendUserVerificationCodeOut, error)
	RegisterUser(r *http.Request) (*emptypb.Empty, error)
	LoginV2(r *http.Request) (*authproto.LoginV2Out, error)
	RefreshAccessToken(r *http.Request) (*authproto.RefreshAccessTokenOut, error)
}
