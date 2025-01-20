package model

type SearchUser struct {
	Nickname   string `json:"nickname"`
	Uuid       string `json:"uuid"`
	AvatarLink string `json:"avatar_link"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	IsFriend   bool   `json:"is_friend"`
}

type SearchUsersOut struct {
	Users []SearchUser `json:"users"`
	Total int64        `json:"total"`
}

type SearchSociety struct {
	Name       string `json:"name"`
	AvatarLink string `json:"avatar_link"`
	SocietyId  int64  `json:"society_id"`
	IsMember   bool   `json:"is_member"`
	IsPrivate  bool   `json:"is_private"`
}

type SearchSocietyOut struct {
	Societies []SearchSociety `json:"societies"`
	Total     int64           `json:"total"`
}
