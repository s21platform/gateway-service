//go:generate mockgen -destination=mock_contract_test.go -package=${GOPACKAGE} -source=contract.go

package api

import (
	"net/http"

	advert "github.com/s21platform/advert-proto/advert-proto"
	avatar "github.com/s21platform/avatar-proto/avatar-proto"
	chat "github.com/s21platform/chat-proto/chat-proto"
	feed "github.com/s21platform/feed-proto/feed-proto"
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
	GetOptionRequests(r *http.Request) (model.OptionRequestsList, error)
}

type SocietyService interface {
	CreateSociety(r *http.Request) (*societyproto.SetSocietyOut, error)
	GetSocietyInfo(r *http.Request) (*model.SocietyInfo, error)
	UpdateSociety(r *http.Request) error
}

type SearchService interface {
	GetUsersWithLimit(r *http.Request) (model.SearchUsersOut, error)
	GetSocietyWithLimit(r *http.Request) (model.SearchSocietyOut, error)
}

type ChatService interface {
	GetChats(r *http.Request) (*chat.GetChatsOut, error)
	CreatePrivateChat(r *http.Request) (*chat.CreatePrivateChatOut, error)
	GetPrivateRecentMessages(r *http.Request) (*chat.GetPrivateRecentMessagesOut, error)
}

type AdvertService interface {
	GetAdverts(r *http.Request) (*advert.GetAdvertsOut, error)
	CreateAdvert(r *http.Request) (*advert.AdvertEmpty, error)
	CancelAdvert(r *http.Request) (*advert.AdvertEmpty, error)
	RestoreAdvert(r *http.Request) (*advert.AdvertEmpty, error)
}

type FeedService interface {
	CreateUserPost(r *http.Request) (*feed.CreateUserPostOut, error)
}
