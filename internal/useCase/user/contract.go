package user

import (
	"context"

	userproto "github.com/s21platform/user-proto/user-proto"
)

type UserClient interface {
	GetInfo(ctx context.Context, uuid string) (*userproto.GetUserInfoByUUIDOut, error)
}
