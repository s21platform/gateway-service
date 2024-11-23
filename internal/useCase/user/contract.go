package user

import (
	"context"

	"github.com/s21platform/gateway-service/internal/model"

	userproto "github.com/s21platform/user-proto/user-proto"
)

type UserClient interface {
	GetInfo(ctx context.Context, uuid string) (*userproto.GetUserInfoByUUIDOut, error)
	UpdateProfile(ctx context.Context, data model.ProfileData) (*userproto.UpdateProfileOut, error)
}
