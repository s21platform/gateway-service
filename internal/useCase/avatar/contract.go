package avatar

import (
	"context"
	"mime/multipart"

	"github.com/s21platform/avatar-service/pkg/avatar"
)

type AvatarClient interface {
	SetUserAvatar(ctx context.Context, filename string, file multipart.File) (*avatar.SetUserAvatarOut, error)
	GetAllUserAvatars(ctx context.Context) (*avatar.GetAllUserAvatarsOut, error)
	DeleteUserAvatar(ctx context.Context, id int32) (*avatar.Avatar, error)

	SetSocietyAvatar(ctx context.Context, filename string, file multipart.File, uuid string) (*avatar.SetSocietyAvatarOut, error)
	GetAllSocietyAvatars(ctx context.Context, uuid string) (*avatar.GetAllSocietyAvatarsOut, error)
	DeleteSocietyAvatar(ctx context.Context, id int32) (*avatar.Avatar, error)
}
