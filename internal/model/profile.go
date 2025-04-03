package model

import (
	"log"
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

	// Check and remove "@" from Telegram username if present
	if len(pd.Telegram) > 0 && pd.Telegram[0] == '@' {
		log.Printf("Telegram username changed from %s to %s", pd.Telegram, pd.Telegram[1:])
		pd.Telegram = pd.Telegram[1:]
	}

	// Преобразование OS ID в nil, если значение равно 0
	var osId *int64
	if pd.Os.Id != 0 {
		osId = &pd.Os.Id
	}

	return &userproto.UpdateProfileIn{
		Name:     pd.Name,
		Birthday: birthday,
		Telegram: pd.Telegram,
		Github:   pd.Git,
		OsId:     *osId,
	}
}
