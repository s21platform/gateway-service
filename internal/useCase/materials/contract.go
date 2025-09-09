package materials

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/s21platform/gateway-service/internal/model"
	materialsproto "github.com/s21platform/materials-service/pkg/materials"
)

type MaterialsClient interface {
	GetAllMaterials(ctx context.Context) (*materialsproto.GetAllMaterialsOut, error)
	EditMaterial(ctx context.Context, req *model.EditMaterialRequest) (*materialsproto.EditMaterialOut, error)
	DeleteMaterial(ctx context.Context, req *model.DeleteMaterialRequest) (*emptypb.Empty, error)
}
