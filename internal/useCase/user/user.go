package user

import (
	"context"

	"github.com/s21platform/gateway-service/internal/config"

	user_proto "github.com/s21platform/user-proto/user-proto"
)

type Usecase struct {
	uC UserClient
}

func New(uC UserClient) *Usecase {
	return &Usecase{uC: uC}
}

func (u *Usecase) GetInfoByUUID(ctx context.Context) (*user_proto.GetUserInfoByUUIDOut, error) {
	uuid := ctx.Value(config.KeyUUID).(string)
	resp, err := u.uC.GetInfo(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
