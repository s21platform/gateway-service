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
	var requestData model.EditMaterialRequest

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode request body: %v", err)
	}
	defer r.Body.Close()

	resp, err := u.mC.EditMaterial(r.Context(), &requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to edit material in usecase: %v", err)
	}

	material := &model.Material{
		UUID:            resp.Material.Uuid,
		OwnerUUID:       resp.Material.OwnerUuid,
		Title:           resp.Material.Title,
		CoverImageURL:   resp.Material.CoverImageUrl,
		Description:     resp.Material.Description,
		Content:         resp.Material.Content,
		ReadTimeMinutes: resp.Material.ReadTimeMinutes,
		Status:          resp.Material.Status,
		CreatedAt:       resp.Material.CreatedAt.AsTime(),
		EditedAt:        resp.Material.EditedAt.AsTime(),
		PublishedAt:     resp.Material.PublishedAt.AsTime(),
		ArchivedAt:      resp.Material.ArchivedAt.AsTime(),
		DeletedAt:       resp.Material.DeletedAt.AsTime(),
		LikesCount:      resp.Material.LikesCount,
	}

	return material, nil
}

func (uc *UseCase) GetAllMaterialsList(r *http.Request) (*model.MaterialList, error) {
	var materialList model.MaterialList

	resp, err := uc.mC.GetAllMaterials(r.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get materials: %w", err)
	}

	materialList.ToDTO(resp.GetMaterialList())
	return &materialList, nil
}

func (uс *UseCase) DeleteMaterial(r *http.Request) error {
	materialUuid := r.URL.Query().Get("id")

	_, err := uс.mC.DeleteMaterial(r.Context(), materialUuid)
	if err != nil {
		return fmt.Errorf("failed to delete material in usecase: %v", err)
	}

	return nil
}

func (uc *UseCase) ArchiveMaterial(r *http.Request) error {
	materialUuid := r.URL.Query().Get("id")

	_, err := uc.mC.ArchiveMaterial(r.Context(), materialUuid)
	if err != nil {
		return fmt.Errorf("failed to archive material in usecase: %v", err)
	}

	return nil
}
