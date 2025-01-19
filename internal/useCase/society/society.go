package society

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	societyproto "github.com/s21platform/society-proto/society-proto"
)

type UseCase struct {
	sC SocietyClient
}

func New(sC SocietyClient) *UseCase {
	return &UseCase{sC: sC}
}

type RequestData struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	IsPrivate     bool   `json:"is_private"`
	DirectionId   int64  `json:"direction_id"`
	AccessLevelId int64  `json:"access_level_id"`
}

type SocietyId struct {
	Id int64 `json:"id"`
}

type Uuid struct {
	Uuid string `json:"uuid"`
}

func (u *UseCase) CreateSociety(r *http.Request) (*societyproto.SetSocietyOut, error) {
	requestData := RequestData{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}
	defer r.Body.Close()

	if len(body) == 0 {
		return nil, fmt.Errorf("request body is empty")
	}

	if err := json.Unmarshal(body, &requestData); err != nil {
		return nil, fmt.Errorf("failed to decode request body: %w", err)
	}

	resp, err := u.sC.CreateSociety(r.Context(), &requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to create society: %v", err)
	}
	return resp, nil
}

func (u *UseCase) GetAccessLevel(r *http.Request) (*societyproto.GetAccessLevelOut, error) {
	resp, err := u.sC.GetAccessLevel(r.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get access level: %v", err)
	}
	return resp, nil
}

func (u *UseCase) GetSocietyInfo(r *http.Request) (*societyproto.GetSocietyInfoOut, error) {
	id := SocietyId{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	if err := json.Unmarshal(body, &id); err != nil {
		return nil, fmt.Errorf("failed to decode request body: %w", err)
	}

	resp, err := u.sC.GetSocietyInfo(r.Context(), id.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get access level: %v", err)
	}

	return resp, nil
}

func (u *UseCase) SubscribeToSociety(r *http.Request) (*societyproto.SubscribeToSocietyOut, error) {
	id := SocietyId{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	if err := json.Unmarshal(body, &id); err != nil {
		return nil, fmt.Errorf("failed to decode request body: %w", err)
	}

	resp, err := u.sC.SubscribeToSociety(r.Context(), id.Id)
	if err != nil {
		return nil, fmt.Errorf("failed subscribe to society: %v", err)
	}

	return resp, nil
}

func (u *UseCase) GetPermission(r *http.Request) (*societyproto.GetPermissionsOut, error) {
	resp, err := u.sC.GetPermission(r.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get permission: %v", err)
	}
	return resp, nil
}

func (u *UseCase) UnsubscribeFromSociety(r *http.Request) (*societyproto.UnsubscribeFromSocietyOut, error) {
	id := SocietyId{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	if err := json.Unmarshal(body, &id); err != nil {
		return nil, fmt.Errorf("failed to decode request body: %w", err)
	}

	resp, err := u.sC.UnsubscribeFromSociety(r.Context(), id.Id)
	if err != nil {
		return nil, fmt.Errorf("failed subscribe to society: %v", err)
	}

	return resp, nil
}

func (u *UseCase) GetSocietiesForUser(r *http.Request) (*societyproto.GetSocietiesForUserOut, error) {
	uuid := Uuid{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	if err := json.Unmarshal(body, &uuid); err != nil {
		return nil, fmt.Errorf("failed to decode request body: %w", err)
	}

	resp, err := u.sC.GetSocietiesForUser(r.Context(), uuid.Uuid)
	if err != nil {
		return nil, fmt.Errorf("failed get society for user: %v", err)
	}

	return resp, nil
}
