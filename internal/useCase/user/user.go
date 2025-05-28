package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/s21platform/user-service/pkg/user"

	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/model"
)

type Usecase struct {
	uC UserClient
}

func New(uC UserClient) *Usecase {
	return &Usecase{uC: uC}
}

func (u *Usecase) GetInfoByUUID(r *http.Request) (*user.GetUserInfoByUUIDOut, error) {
	uuid := r.Context().Value(config.KeyUUID).(string)
	resp, err := u.uC.GetInfo(r.Context(), uuid)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *Usecase) GetPeerInfo(r *http.Request) (*user.GetUserInfoByUUIDOut, error) {
	uuid := r.PathValue("uuid")
	if uuid == "" {
		return nil, fmt.Errorf("uuid is empty")
	}
	resp, err := u.uC.GetInfo(r.Context(), uuid)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *Usecase) UpdateProfileInfo(r *http.Request) (*user.UpdateProfileOut, error) {
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

func (u *Usecase) CreatePost(r *http.Request) (*user.CreatePostOut, error) {
	content := r.URL.Query().Get("content")

	resp, err := u.uC.CreatePost(r.Context(), content)
	if err != nil {
		return nil, fmt.Errorf("failed to create post in usecase: %v", err)
	}
	return resp, nil
}