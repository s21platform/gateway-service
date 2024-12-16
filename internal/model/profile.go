package model

import (
	"time"

	userproto "github.com/s21platform/user-proto/user-proto"
)

type OS struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ProfileData struct {
	Name      string     `json:"name"`
	Birthdate *time.Time `json:"birthdate"`
	Telegram  string     `json:"telegram"`
	Git       string     `json:"git"`
	Os        OS         `json:"os"`
}

func (pd *ProfileData) FromDTO() *userproto.UpdateProfileIn {
	var birthday string
	if pd.Birthdate != nil {
		birthday = pd.Birthdate.Format(time.RFC3339)
	}

	//TODO: telegram

	return &userproto.UpdateProfileIn{
		Name:     pd.Name,
		Birthday: birthday,
		Telegram: pd.Telegram,
		Github:   pd.Git,
		OsId:     pd.Os.Id,
	}
}
