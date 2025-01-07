package avatar

import (
	"context"
	"mime/multipart"

	avatar "github.com/s21platform/avatar-proto/avatar-proto"
)

type AvatarClient interface {
	SetUserAvatar(ctx context.Context, filename string, file multipart.File, uuid string) (*avatar.SetUserAvatarOut, error)
	GetAllUserAvatars(ctx context.Context, uuid string) (*avatar.GetAllUserAvatarsOut, error)
	DeleteUserAvatar(ctx context.Context, id int32) (*avatar.Avatar, error)

	SetSocietyAvatar(ctx context.Context, filename string, file multipart.File, uuid string) (*avatar.SetSocietyAvatarOut, error)
	GetAllSocietyAvatars(ctx context.Context, uuid string) (*avatar.GetAllSocietyAvatarsOut, error)
	DeleteSocietyAvatar(ctx context.Context, id int32) (*avatar.Avatar, error)
}
