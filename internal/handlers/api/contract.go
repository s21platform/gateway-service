package api

import (
	"net/http"

	"github.com/s21platform/search-proto/search"

	avatar "github.com/s21platform/avatar-proto/avatar-proto"
	friends "github.com/s21platform/friends-proto/friends-proto"
	notificationproto "github.com/s21platform/notification-proto/notification-proto"
	optionhub "github.com/s21platform/optionhub-proto/optionhub-proto"
	societyproto "github.com/s21platform/society-proto/society-proto"
	userproto "github.com/s21platform/user-proto/user-proto"

	"github.com/s21platform/gateway-service/internal/model"
)

type UserService interface {
	GetInfoByUUID(r *http.Request) (*userproto.GetUserInfoByUUIDOut, error)
	UpdateProfileInfo(r *http.Request) (*userproto.UpdateProfileOut, error)
	GetPeerInfo(r *http.Request) (*userproto.GetUserInfoByUUIDOut, error)
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
	SetFriends(r *http.Request) (*friends.SetFriendsOut, error)
	RemoveFriends(r *http.Request) (*friends.RemoveFriendsOut, error)
}

type OptionService interface {
	GetOsList(r *http.Request) (*model.OptionsStruct, error)
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
	GetSocietyInfo(r *http.Request) (*societyproto.GetSocietyInfoOut, error)
}

type SearchService interface {
	GetUsersWithLimit(r *http.Request) (*search.GetUserWithLimitOut, error)
}
