package community

import (
	"context"

	"github.com/s21platform/gateway-service/internal/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommunityClient interface {
	SendEduLinkingCode(ctx context.Context, in *model.SendEduLinkingCodeRequestData) (*emptypb.Empty, error)
}
