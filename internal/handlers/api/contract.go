package api

import (
	"context"
	"net/http"

	friends "github.com/s21platform/friends-proto/friends-proto"

	notificationproto "github.com/s21platform/notification-proto/notification-proto"

	avatar "github.com/s21platform/avatar-proto/avatar-proto"
	optionhub "github.com/s21platform/optionhub-proto/optionhub-proto"

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

type NotificationService interface {
	GetCountNotification(r *http.Request) (*notificationproto.NotificationCountOut, error)
	GetNotification(r *http.Request) (*notificationproto.NotificationOut, error)
}

type FriendsService interface {
	GetCountFriends(r *http.Request) (*friends.GetCountFriendsOut, error)
}

type OptionService interface {
	GetOsList(r *http.Request) (*optionhub.GetAllOut, error)
}
