package society

import (
	"context"

	societyproto "github.com/s21platform/society-proto/society-proto"
)

type SocietyClient interface {
	CreateSociety(ctx context.Context, req *RequestData) (*societyproto.SetSocietyOut, error)
}
