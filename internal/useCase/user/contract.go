package user

import (
	"context"

	"github.com/s21platform/user-service/pkg/user"

	"github.com/s21platform/gateway-service/internal/model"
)

type UserClient interface {
	GetInfo(ctx context.Context, uuid string) (*user.GetUserInfoByUUIDOut, error)
	UpdateProfile(ctx context.Context, data model.ProfileData) (*user.UpdateProfileOut, error)
	CreatePost(ctx context.Context, content string) (*user.CreatePostOut, error)
}
