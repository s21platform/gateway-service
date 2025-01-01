package society

import (
	"context"

	societyproto "github.com/s21platform/society-proto/society-proto"
)

type SocietyClient interface {
	CreateSociety(ctx context.Context, req *RequestData) (*societyproto.SetSocietyOut, error)
	GetAccessLevel(ctx context.Context) (*societyproto.GetAccessLevelOut, error)
	GetSocietyInfo(ctx context.Context, id int64) (*societyproto.GetSocietyInfoOut, error)
	SubscribeToSociety(ctx context.Context, id int64) (*societyproto.SubscribeToSocietyOut, error)
	GetPermission(ctx context.Context) (*societyproto.GetPermissionsOut, error)
	UnsubscribeFromSociety(ctx context.Context, id int64) (*societyproto.UnsubscribeFromSocietyOut, error)
}
