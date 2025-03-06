package model

import "time"

type AdvertRequestData struct {
	OwnerUUID   string     `json:"uuid"`
	TextContent string     `json:"text"`
	UserFilter  UserFilter `json:"user"`
	ExpiredAt   time.Time  `json:"expires_at"`
}

type UserFilter struct {
	Os []int64 `json:"os,omitempty"`
}

type RestoreAdvertRequestData struct {
	AdvertId int64 `json:"id"`
}
