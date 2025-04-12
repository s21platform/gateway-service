package model

type MarkNotificationsAsReadRequest struct {
	Data struct {
		IDs []int64 `json:"ids"`
	} `json:"data"`
}
