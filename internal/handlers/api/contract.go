//go:generate mockgen -destination=mock_contract_test.go -package=${GOPACKAGE} -source=contract.go

package api

import (
	"net/http"

	chat "github.com/s21platform/chat-proto/chat-proto"

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
	UploadUserAvatar(r *http.Request) (*avatar.SetUserAvatarOut, error)
	GetUserAvatarsList(r *http.Request) (*avatar.GetAllUserAvatarsOut, error)
	RemoveUserAvatar(r *http.Request) (*avatar.Avatar, error)

	UploadSocietyAvatar(r *http.Request) (*avatar.SetSocietyAvatarOut, error)
	GetSocietyAvatarsList(r *http.Request) (*avatar.GetAllSocietyAvatarsOut, error)
	RemoveSocietyAvatar(r *http.Request) (*avatar.Avatar, error)
}

type NotificationService interface {
	GetCountNotification(r *http.Request) (*notificationproto.NotificationCountOut, error)
	GetNotification(r *http.Request) (*notificationproto.NotificationOut, error)
}

type FriendsService interface {
	GetCountFriends(r *http.Request) (*friends.GetCountFriendsOut, error)
	SetFriends(r *http.Request) (*friends.SetFriendsOut, error)
	RemoveFriends(r *http.Request) (*friends.RemoveFriendsOut, error)
	CheckSubscribe(r *http.Request) (*model.CheckSubscribe, error)
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
	RemoveSociety(r *http.Request) (*societyproto.EmptySociety, error)
	//GetAccessLevel(r *http.Request) (*societyproto.GetAccessLevelOut, error)
	//GetSocietyInfo(r *http.Request) (*societyproto.GetSocietyInfoOut, error)
	//SubscribeToSociety(r *http.Request) (*societyproto.SubscribeToSocietyOut, error)
	//GetPermission(r *http.Request) (*societyproto.GetPermissionsOut, error)
	//UnsubscribeFromSociety(r *http.Request) (*societyproto.UnsubscribeFromSocietyOut, error)
	//GetSocietiesForUser(r *http.Request) (*societyproto.GetSocietiesForUserOut, error)
}

type SearchService interface {
	GetUsersWithLimit(r *http.Request) (model.SearchUsersOut, error)
	GetSocietyWithLimit(r *http.Request) (model.SearchSocietyOut, error)
}

type ChatService interface {
	GetRecentMessages(r *http.Request) (*chat.GetRecentMessagesOut, error)
}
