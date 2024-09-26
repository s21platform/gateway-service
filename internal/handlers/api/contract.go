package api

import (
	"context"
	user_proto "github.com/s21platform/user-proto/user-proto"
)

type UserService interface {
	GetInfoByUUID(ctx context.Context) (*user_proto.GetUserInfoByUUIDOut, error)
}
