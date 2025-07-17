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
	FormatID         int64   `json:"format_id"`
	PostPermissionID int64   `json:"post_permission_id"`
	IsSearch         bool    `json:"is_search"`
	Tags             []int64 `json:"tags"`
}

type SocietyInfo struct {
	SocietyUUID      string  `json:"society_id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	OwnerUUID        string  `json:"owner_uuid"`
	PhotoURL         string  `json:"photo_url"`
	FormatID         int64   `json:"format_id"`
	PostPermissionID int64   `json:"post_permission_id"`
	IsSearch         bool    `json:"is_search"`
	CountSubscribe   int64   `json:"count_subscribe"`
	Tags             []int64 `json:"tags"`
	CanEditSociety   bool    `json:"can_edit_society"`
}

type SocietyId struct {
	Id string `json:"id"`
}
