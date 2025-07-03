package avatar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/s21platform/avatar-service/pkg/avatar"
)

type Usecase struct {
	aC AvatarClient
}

func New(aC AvatarClient) *Usecase {
	return &Usecase{aC: aC}
}

func (uc *Usecase) UploadUserAvatar(r *http.Request) (*avatar.SetUserAvatarOut, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, fmt.Errorf("failed to parse multipart form: %w", err)
	}

	file, _, err := r.FormFile("avatar")
	if err != nil {
		return nil, fmt.Errorf("failed to get file from form: %w", err)
	}
	defer file.Close()

	filename := r.FormValue("filename")

	resp, err := uc.aC.SetUserAvatar(r.Context(), filename, file)
	if err != nil {
		return nil, fmt.Errorf("failed to set avatar: %w", err)
	}

	return resp, nil
}

func (uc *Usecase) GetUserAvatarsList(r *http.Request) (*avatar.GetAllUserAvatarsOut, error) {
	resp, err := uc.aC.GetAllUserAvatars(r.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get avatars: %w", err)
	}

	return resp, nil
}

func (uc *Usecase) RemoveUserAvatar(r *http.Request) (*avatar.Avatar, error) {
	id, err := getAvatarId(r)
	if err != nil {
		return nil, fmt.Errorf("failed to get avatar id: %w", err)
	}

	resp, err := uc.aC.DeleteUserAvatar(r.Context(), id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete avatar: %w", err)
	}

	return resp, nil
}

func (uc *Usecase) UploadSocietyAvatar(r *http.Request) (*avatar.SetSocietyAvatarOut, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, fmt.Errorf("failed to parse multipart form: %w", err)
	}

	file, _, err := r.FormFile("avatar")
	if err != nil {
		return nil, fmt.Errorf("failed to get file from form: %w", err)
	}
	defer file.Close()

	uuid := r.FormValue("societyUUID")
	filename := r.FormValue("filename")

	resp, err := uc.aC.SetSocietyAvatar(r.Context(), filename, file, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to set avatar: %w", err)
	}

	return resp, nil
}

func (uc *Usecase) GetSocietyAvatarsList(r *http.Request) (*avatar.GetAllSocietyAvatarsOut, error) {
	uuid := r.URL.Query().Get("societyUUID")
	if uuid == "" {
		return nil, fmt.Errorf("failed to no society UUID in request")
	}

	resp, err := uc.aC.GetAllSocietyAvatars(r.Context(), uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get avatars: %w", err)
	}

	return resp, nil
}

func (uc *Usecase) RemoveSocietyAvatar(r *http.Request) (*avatar.Avatar, error) {
	id, err := getAvatarId(r)
	if err != nil {
		return nil, fmt.Errorf("failed to get avatar id: %w", err)
	}

	resp, err := uc.aC.DeleteSocietyAvatar(r.Context(), id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete avatar: %w", err)
	}

	return resp, nil
}

func getAvatarId(r *http.Request) (int32, error) {
	var requestData struct {
		ID int32 `json:"id"`
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read request body: %w", err)
	}
	defer r.Body.Close()

	if len(body) == 0 {
		return 0, fmt.Errorf("request body is empty")
	}

	if err = json.Unmarshal(body, &requestData); err != nil {
		return 0, fmt.Errorf("failed to decode request body: %w", err)
	}

	return requestData.ID, nil
}
