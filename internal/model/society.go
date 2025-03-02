package model

type JoinStatus struct {
	Success bool `json:"success"`
}

type RequestData struct {
	Name             string `json:"name"`
	FormatID         int64  `json:"format_id"`
	PostPermissionID int64  `json:"post_permission_id"`
	IsSearch         bool   `json:"is_search"`
}

type SocietyUpdate struct {
	SocietyUUID      string  `json:"society_id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	PhotoURL         string  `json:"photo_url"`
	FormatID         int64   `json:"format_id"`
	PostPermissionID int64   `json:"post_permission_id"`
	IsSearch         bool    `json:"is_search"`
	Tags             []int64 `json:"tags"`
}
