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
	FullName  string    `json:"fullName"`
	BirthDate time.Time `json:"birthDate"`
	Telegram  string    `json:"telegram"`
	GitLink   string    `json:"gitLink"`
	Os        OS        `json:"os"`
}

func (pd *ProfileData) FromDTO() *userproto.UpdateProfileIn {
	return &userproto.UpdateProfileIn{
		Name:     pd.FullName,
		Birthday: pd.BirthDate.Format(time.RFC3339),
		Telegram: pd.Telegram,
		Github:   pd.GitLink,
		OsId:     pd.Os.Id,
	}
}
