package materials

import (
	"context"

	"github.com/s21platform/materials-service/pkg/materials"
)

type MaterialsClient interface {
	EditMaterial(ctx context.Context, uuid string) (*materials.EditMaterialOut, error)
}
