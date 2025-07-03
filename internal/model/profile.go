package model

import (
	"log"
	"time"

	"github.com/s21platform/user-service/pkg/user"
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

func (pd *ProfileData) FromDTO() *user.UpdateProfileIn {
	var birthday string
	if pd.Birthdate != nil {
		birthday = pd.Birthdate.Format(time.RFC3339)
	}

	// Check and remove "@" from Telegram username if present
	if len(pd.Telegram) > 0 && pd.Telegram[0] == '@' {
		log.Printf("Telegram username changed from %s to %s", pd.Telegram, pd.Telegram[1:])
		pd.Telegram = pd.Telegram[1:]
	}

	return &user.UpdateProfileIn{
		Name:     pd.Name,
		Birthday: birthday,
		Telegram: pd.Telegram,
		Github:   pd.Git,
		OsId:     pd.Os.Id,
	}
}
