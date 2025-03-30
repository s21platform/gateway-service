package model

import (
	"time"

	v1 "github.com/s21platform/optionhub-proto/optionhub/v1"
)

type OptionsStruct struct {
	Options []Option `json:"options"`
}

type Option struct {
	Id    int64  `json:"id"`
	Label string `json:"label"`
}

type OptionRequest struct {
	ID             int64     `json:"id"`
	AttributeID    int64     `json:"attribute_id"`
	AttributeValue string    `json:"attribute_value"`
	Value          string    `json:"value"`
	UserUuid       string    `json:"user_uuid"`
	CreatedAt      time.Time `json:"created_at"`
}

type OptionRequestsList []OptionRequest

func (o *OptionRequestsList) FromDTO(obj *v1.GetOptionRequestsOut) {
	tmp := OptionRequestsList{}
	for _, or := range obj.OptionRequestItem {
		tmp = append(tmp, OptionRequest{
			ID:             or.OptionRequestId,
			AttributeID:    or.AttributeId,
			AttributeValue: or.AttributeValue,
			Value:          or.OptionRequestValue,
			UserUuid:       or.UserUuid,
			CreatedAt:      or.CreatedAt.AsTime(),
		})
	}
	*o = tmp
}
