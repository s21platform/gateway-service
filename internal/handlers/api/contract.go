package api

import (
	"context"
	"net/http"

	avatar "github.com/s21platform/avatar-proto/avatar-proto"

	user_proto "github.com/s21platform/user-proto/user-proto"
)

type UserService interface {
	GetInfoByUUID(ctx context.Context) (*user_proto.GetUserInfoByUUIDOut, error)
}

type AvatarService interface {
	UploadAvatar(r *http.Request) (*avatar.SetAvatarOut, error)
}
