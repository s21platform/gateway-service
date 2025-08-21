package materials

import (
	"context"

	"github.com/s21platform/gateway-service/internal/model"
	materialsproto "github.com/s21platform/materials-service/pkg/materials"
)

type MaterialsClient interface {
	EditMaterial(ctx context.Context, req *model.EditMaterial) (*materialsproto.EditMaterialOut, error)
}
