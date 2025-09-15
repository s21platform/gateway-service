package model

type SendEduLinkingCodeRequestData struct {
	Login string `json:"login"`
}

type ValidateCode struct {
	Login string `json:"login"`
	Code  int64  `json:"code"`
}
