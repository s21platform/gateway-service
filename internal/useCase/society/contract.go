package society

import (
	"context"

	"github.com/s21platform/gateway-service/internal/model"

	societyproto "github.com/s21platform/society-proto/society-proto"
)

type SocietyClient interface {
	CreateSociety(ctx context.Context, req *model.RequestData) (*societyproto.SetSocietyOut, error)
	GetSocietyInfo(ctx context.Context, societyInfo string) (*societyproto.GetSocietyInfoOut, error)
	UpdateSociety(ctx context.Context, req *model.SocietyUpdate) error
}
