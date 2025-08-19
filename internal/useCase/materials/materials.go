package materials

import (
	"fmt"
	"net/http"

	"github.com/s21platform/materials-service/pkg/materials"
)

type UseCase struct {
	mC MaterialsClient
}

func New(mC MaterialsClient) *UseCase {
	return &UseCase{mC: mC}
}

func (uc *UseCase) GetAllMaterialsList(r *http.Request) (*materials.GetAllMaterialsOut, error) {
	resp, err := uc.mC.GetAllMaterials(r.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get avatars: %w", err)
	}

	return resp, nil
}
