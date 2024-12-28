package model

type SearchUser struct {
	Nickname   string `json:"nickname"`
	Uuid       string `json:"uuid"`
	AvatarLink string `json:"avatar_link"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
}

type SearchUsersOut struct {
	Users []SearchUser `json:"users"`
	Total int64        `json:"total"`
}
