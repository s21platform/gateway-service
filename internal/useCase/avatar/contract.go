package avatar

import (
	"context"
	"mime/multipart"

	avatar "github.com/s21platform/avatar-proto/avatar-proto"
)

type AvatarClient interface {
	SetAvatar(ctx context.Context, filename string, file multipart.File, uuid string) (*avatar.SetAvatarOut, error)
	GetAllAvatars(ctx context.Context, uuid string) (*avatar.GetAllAvatarsOut, error)
	DeleteAvatar(ctx context.Context, id int32) (*avatar.Avatar, error)
}
