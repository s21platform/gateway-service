package materials

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/s21platform/materials-service/pkg/materials"

	"github.com/s21platform/gateway-service/internal/model"
)

type UseCase struct {
	mC MaterialsClient
}

func New(mC MaterialsClient) *UseCase {
	return &UseCase{mC: mC}
}

func (u *UseCase) EditMaterial(r *http.Request) (*materials.EditMaterialOut, error) {
	var requestData model.EditMaterial

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode request body: %v", err)
	}
	defer r.Body.Close()

	resp, err := u.mC.EditMaterial(r.Context(), requestData.MaterialUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to edit material in usecase: %v", err)
	}

	return resp, nil
}
