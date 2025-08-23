package community

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/s21platform/gateway-service/internal/model"
)

type CommunityClient interface {
	SendEduLinkingCode(ctx context.Context, in *model.SendEduLinkingCodeRequestData) (*emptypb.Empty, error)
}
