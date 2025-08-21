package materials

import (
	"context"

	"github.com/s21platform/gateway-service/internal/model"
)

type MaterialsClient interface {
	GetAllMaterials(ctx context.Context) (*model.MaterialList, error)
}
