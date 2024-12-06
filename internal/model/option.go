package model

type OptionsStruct struct {
	Options []Option `json:"options"`
}

type Option struct {
	Id    int64  `json:"id"`
	Label string `json:"label"`
}
