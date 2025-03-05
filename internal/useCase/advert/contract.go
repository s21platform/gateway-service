package advert

import (
	"context"

	advert "github.com/s21platform/advert-proto/advert-proto"

	"github.com/s21platform/gateway-service/internal/model"
)

type AdvertClient interface {
	GetAdverts(ctx context.Context, uuid string) (*advert.GetAdvertsOut, error)
	CreateAdvert(ctx context.Context, req *model.CreateAdvertRequestData) (*advert.AdvertEmpty, error)
	CancelAdvert(ctx context.Context, req *model.CancelAdvertRequestData) (*advert.AdvertEmpty, error)
}
