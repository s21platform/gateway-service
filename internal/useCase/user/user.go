package user

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/s21platform/gateway-service/internal/model"

	"github.com/s21platform/gateway-service/internal/config"

	userproto "github.com/s21platform/user-proto/user-proto"
)

type Usecase struct {
	uC UserClient
}

func New(uC UserClient) *Usecase {
	return &Usecase{uC: uC}
}

func (u *Usecase) GetInfoByUUID(ctx context.Context) (*userproto.GetUserInfoByUUIDOut, error) {
	uuid := ctx.Value(config.KeyUUID).(string)
	resp, err := u.uC.GetInfo(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *Usecase) UpdateProfileInfo(r *http.Request) (*userproto.UpdateProfileOut, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body for update profile: %w", err)
	}
	defer r.Body.Close()

	var data model.ProfileData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body for update profile: %w", err)
	}
	fmt.Println(data)

	resp, err := u.uC.UpdateProfile(r.Context(), data)
	if err != nil {
		return nil, fmt.Errorf("failed to update profile: %w", err)
	}

	return resp, nil
}
