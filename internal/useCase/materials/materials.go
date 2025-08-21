package materials

import (
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

func (uc *UseCase) GetAllMaterialsList(r *http.Request) (*model.MaterialList, error) {
	resp, err := uc.mC.GetAllMaterials(r.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get materials: %w", err)
	}
	return resp, nil
}
