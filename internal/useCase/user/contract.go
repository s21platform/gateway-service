package user

import (
	"context"

	userproto "github.com/s21platform/user-proto/user-proto"

	"github.com/s21platform/gateway-service/internal/model"
)

type UserClient interface {
	GetInfo(ctx context.Context, uuid string) (*userproto.GetUserInfoByUUIDOut, error)
	UpdateProfile(ctx context.Context, data model.ProfileData) (*userproto.UpdateProfileOut, error)
}
