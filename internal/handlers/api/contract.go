package api

import (
	"context"
	"net/http"

	societyproto "github.com/s21platform/society-proto/society-proto"

	friends "github.com/s21platform/friends-proto/friends-proto"

	notificationproto "github.com/s21platform/notification-proto/notification-proto"

	avatar "github.com/s21platform/avatar-proto/avatar-proto"
	optionhub "github.com/s21platform/optionhub-proto/optionhub-proto"

	userproto "github.com/s21platform/user-proto/user-proto"
)

type UserService interface {
	GetInfoByUUID(ctx context.Context) (*userproto.GetUserInfoByUUIDOut, error)
	UpdateProfileInfo(r *http.Request) (*userproto.UpdateProfileOut, error)
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
	GetOsList(r *http.Request) (*optionhub.GetByNameOut, error)
	GetWorkPlaceList(r *http.Request) (*optionhub.GetByNameOut, error)
	GetStudyPlaceList(r *http.Request) (*optionhub.GetByNameOut, error)
	GetHobbyList(r *http.Request) (*optionhub.GetByNameOut, error)
	GetSkillList(r *http.Request) (*optionhub.GetByNameOut, error)
	GetCityList(r *http.Request) (*optionhub.GetByNameOut, error)
	GetSocietyDirectionList(r *http.Request) (*optionhub.GetByNameOut, error)
}

type SocietyService interface {
	CreateSociety(r *http.Request) (*societyproto.SetSocietyOut, error)
	GetAccessLevel(r *http.Request) (*societyproto.GetAccessLevelOut, error)
}
