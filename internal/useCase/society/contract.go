package society

import (
	"context"

	society_proto "github.com/s21platform/society-proto/society-proto"
)

type SocietyClient interface {
	CreateSociety(ctx context.Context, req *RequestData) (*society_proto.SetSocietyOut, error)
}
