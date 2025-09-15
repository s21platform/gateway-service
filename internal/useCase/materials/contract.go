package materials

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	materialsproto "github.com/s21platform/materials-service/pkg/materials"

	"github.com/s21platform/gateway-service/internal/model"
)

type MaterialsClient interface {
	GetAllMaterials(ctx context.Context) (*materialsproto.GetAllMaterialsOut, error)
	EditMaterial(ctx context.Context, req *model.EditMaterialRequest) (*materialsproto.EditMaterialOut, error)
	DeleteMaterial(ctx context.Context, materialUuid string) (*emptypb.Empty, error)
	ArchiveMaterial(ctx context.Context, materialUuid string) (*emptypb.Empty, error)
}
