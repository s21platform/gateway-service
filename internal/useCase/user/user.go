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
	var req model.CreatePostRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("failed to decode request body: %v", err)
	}
	defer r.Body.Close()

	resp, err := u.uC.CreatePost(r.Context(), req.Content)
	if err != nil {
		return nil, fmt.Errorf("failed to create post in usecase: %v", err)
	}
	return resp, nil
}

func (u *Usecase) SetUserFriends(r *http.Request) (*user.SetFriendsOut, error) {
	var readPeer struct {
		Peer string `json:"peer"`
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Body: %v", err)
	}
	defer r.Body.Close()

	if len(body) == 0 {
		return nil, fmt.Errorf("failed to request body is empty")
	}
	if err = json.Unmarshal(body, &readPeer); err != nil {
		return nil, fmt.Errorf("failed to json unmarshal: %v", err)
	}

	resp, err := u.uC.SetFriends(r.Context(), &user.SetFriendsIn{Peer: readPeer.Peer})
	if err != nil {
		return nil, fmt.Errorf("failed to user service Set Friends: %v", err)
	}

	return resp, nil
}

func (u *Usecase) RemoveUserFriends(r *http.Request) (*user.RemoveFriendsOut, error) {
	var readPeer struct {
		Peer string `json:"peer"`
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Body: %v", err)
	}
	defer r.Body.Close()
	if len(body) == 0 {
		return nil, fmt.Errorf("failed to request body is empty")
	}
	if err = json.Unmarshal(body, &readPeer); err != nil {
		return nil, fmt.Errorf("failed to json unmarshal: %v", err)
	}

	resp, err := u.uC.RemoveFriends(r.Context(), &user.RemoveFriendsIn{Peer: readPeer.Peer})
	if err != nil {
		return nil, fmt.Errorf("failed to user service RemoveFriends: %v", err)
	}
	return resp, nil
}

func (u *Usecase) GetUserCountFriends(r *http.Request) (*user.GetCountFriendsOut, error) {
	resp, err := u.uC.GetCountFriends(r.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to u.fC.GetCountFriends: %v", err)
	}
	return resp, nil
}
