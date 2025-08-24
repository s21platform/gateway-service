package materials

import (
	"context"

	"github.com/s21platform/materials-service/pkg/materials"
)

type MaterialsClient interface {
	GetAllMaterials(ctx context.Context) (*materials.GetAllMaterialsOut, error)
}
