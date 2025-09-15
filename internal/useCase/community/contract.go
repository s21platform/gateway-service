package community

import (
	"context"
	"github.com/s21platform/community-service/pkg/community"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/s21platform/gateway-service/internal/model"
)

type CommunityClient interface {
	SendEduLinkingCode(ctx context.Context, in *model.SendEduLinkingCodeRequestData) (*emptypb.Empty, error)
	ValidateCode(ctx context.Context, in *model.ValidateCode) (*community.ValidateCodeOut, error)
}
