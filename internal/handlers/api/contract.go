package api

import (
	"context"
	"net/http"

	avatar "github.com/s21platform/avatar-proto/avatar-proto"

	userproto "github.com/s21platform/user-proto/user-proto"
)

type UserService interface {
	GetInfoByUUID(ctx context.Context) (*userproto.GetUserInfoByUUIDOut, error)
}

type AvatarService interface {
	UploadAvatar(r *http.Request) (*avatar.SetAvatarOut, error)
	GetAvatarsList(r *http.Request) (*avatar.GetAllAvatarsOut, error)
	RemoveAvatar(r *http.Request) (*avatar.Avatar, error)
}
