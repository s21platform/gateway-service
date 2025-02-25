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
	Name             string `json:"name"`
	FormatID         int64  `json:"format_id"`
	PostPermissionID int64  `json:"post_permission_id"`
	IsSearch         bool   `json:"is_search"`
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

func (u *UseCase) GetSocietyInfo(r *http.Request) (*societyproto.GetSocietyInfoOut, error) {
	societyUUID := r.URL.Query().Get("society_id")

	res, err := u.sC.GetSocietyInfo(r.Context(), societyUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get society: %v", err)
	}
	return res, nil
}

//
//func (u *UseCase) SubscribeToSociety(r *http.Request) (*societyproto.SubscribeToSocietyOut, error) {
//	strId := r.URL.Query().Get("id")
//	id, err := strconv.Atoi(strId)
//	if err != nil {
//		return nil, fmt.Errorf("failed to convert id to int: %v", err)
//	}
//	resp, err := u.sC.SubscribeToSociety(r.Context(), int64(id))
//	if err != nil {
//		return nil, fmt.Errorf("failed subscribe to society: %v", err)
//	}
//
//	return resp, nil
//}
//
//func (u *UseCase) UnsubscribeFromSociety(r *http.Request) (*societyproto.UnsubscribeFromSocietyOut, error) {
//	id := SocietyId{}
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		return nil, fmt.Errorf("failed to read request body: %w", err)
//	}
//
//	if err := json.Unmarshal(body, &id); err != nil {
//		return nil, fmt.Errorf("failed to decode request body: %w", err)
//	}
//
//	resp, err := u.sC.UnsubscribeFromSociety(r.Context(), id.Id)
//	if err != nil {
//		return nil, fmt.Errorf("failed subscribe to society: %v", err)
//	}
//
//	return resp, nil
//}
//
//func (u *UseCase) GetSocietiesForUser(r *http.Request) (*societyproto.GetSocietiesForUserOut, error) {
//	uuid := r.URL.Query().Get("uuid")
//
//	resp, err := u.sC.GetSocietiesForUser(r.Context(), uuid)
//	if err != nil {
//		return nil, fmt.Errorf("failed get society for user: %v", err)
//	}
//
//	return resp, nil
//}
