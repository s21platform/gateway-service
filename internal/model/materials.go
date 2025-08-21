package model

import "time"

type Material struct {
	UUID            string    `json:"uuid"`
	OwnerUUID       string    `json:"owner_uuid"`
	Title           string    `json:"title"`
	CoverImageURL   string    `json:"cover_image_url"`
	Description     string    `json:"description"`
	Content         string    `json:"content"`
	ReadTimeMinutes int32     `json:"read_time_minutes"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	EditedAt        time.Time `json:"edited_at"`
	PublishedAt     time.Time `json:"published_at"`
	ArchivedAt      time.Time `json:"archived_at"`
	DeletedAt       time.Time `json:"deleted_at"`
	LikesCount      int32     `json:"likes_count"`
}

type EditMaterial struct {
	UUID            string `json:"uuid"`
	Title           string `json:"title"`
	CoverImageURL   string `json:"cover_image_url"`
	Description     string `json:"description"`
	Content         string `json:"content"`
	ReadTimeMinutes int32  `json:"read_time_minutes"`
}
