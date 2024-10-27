package avatar

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/s21platform/gateway-service/internal/config"

	avatar "github.com/s21platform/avatar-proto/avatar-proto"
)

type Usecase struct {
	aC AvatarClient
}

func New(aC AvatarClient) *Usecase {
	return &Usecase{aC: aC}
}

func (uc *Usecase) UploadAvatar(r *http.Request) (*avatar.SetAvatarOut, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, fmt.Errorf("failed to parse multipart form: %w", err)
	}
	file, _, err := r.FormFile("avatar")
	if err != nil {
		return nil, fmt.Errorf("failed to get file from form: %w", err)
	}
	defer file.Close()

	uuid := r.Context().Value(config.KeyUUID).(string)
	filename := r.FormValue("filename")

	resp, err := uc.aC.SetAvatar(r.Context(), filename, file, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to set avatar: %w", err)
	}
	return resp, nil
}

func (uc *Usecase) GetAvatarsList(r *http.Request) (*avatar.GetAllAvatarsOut, error) {
	uuid := r.Context().Value(config.KeyUUID).(string)

	resp, err := uc.aC.GetAllAvatars(r.Context(), uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get avatars: %w", err)
	}

	return resp, nil
}

func (uc *Usecase) RemoveAvatar(r *http.Request) (*avatar.Avatar, error) {
	id, err := getAvatarId(r)
	if err != nil {
		return nil, fmt.Errorf("failed to get avatar id: %w", err)
	}

	resp, err := uc.aC.DeleteAvatar(r.Context(), id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete avatar: %w", err)
	}

	return resp, nil
}

func getAvatarId(r *http.Request) (int32, error) {
	var requestData struct {
		ID int32 `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		return 0, fmt.Errorf("failed to decode request body: %w", err)
	}

	return requestData.ID, nil
}
