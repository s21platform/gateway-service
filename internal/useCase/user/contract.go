package user

import (
	"context"

	"github.com/s21platform/gateway-service/internal/model"
	userproto "github.com/s21platform/user-proto/user-proto"
	user "github.com/s21platform/user-service/pkg/user"
)

type UserClient interface {
	GetInfo(ctx context.Context, uuid string) (*userproto.GetUserInfoByUUIDOut, error)
	UpdateProfile(ctx context.Context, data model.ProfileData) (*userproto.UpdateProfileOut, error)
	SetFriends(ctx context.Context, peer *user.SetFriendsIn) (*user.SetFriendsOut, error)
	RemoveFriends(ctx context.Context, peer *user.RemoveFriendsIn) (*user.RemoveFriendsOut, error)
	GetCountFriends(ctx context.Context) (*user.GetCountFriendsOut, error)
}
