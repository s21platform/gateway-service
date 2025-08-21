package materials

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/s21platform/gateway-service/internal/model"
)

type UseCase struct {
	mC MaterialsClient
}

func New(mC MaterialsClient) *UseCase {
	return &UseCase{mC: mC}
}

func (u *UseCase) EditMaterial(r *http.Request) (*model.Material, error) {
	var requestData model.EditMaterial

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode request body: %v", err)
	}
	defer r.Body.Close()

	resp, err := u.mC.EditMaterial(r.Context(), &requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to edit material in usecase: %v", err)
	}

	mat := resp.GetMaterial()

	material := &model.Material{
		UUID:            mat.Uuid,
		OwnerUUID:       mat.OwnerUuid,
		Title:           mat.Title,
		CoverImageURL:   mat.CoverImageUrl,
		Description:     mat.Description,
		Content:         mat.Content,
		ReadTimeMinutes: mat.ReadTimeMinutes,
		Status:          mat.Status,
		CreatedAt:       mat.CreatedAt.AsTime(),
		EditedAt:        mat.EditedAt.AsTime(),
		PublishedAt:     mat.PublishedAt.AsTime(),
		ArchivedAt:      mat.ArchivedAt.AsTime(),
		DeletedAt:       mat.DeletedAt.AsTime(),
		LikesCount:      mat.LikesCount,
	}

	return material, nil
}
