package materials

import (
	"context"

	materialsproto "github.com/s21platform/materials-service/pkg/materials"

	"github.com/s21platform/gateway-service/internal/model"
)

type MaterialsClient interface {
	EditMaterial(ctx context.Context, req *model.EditMaterialRequest) (*materialsproto.EditMaterialOut, error)
}
