package advert

import (
	"context"

	advert "github.com/s21platform/advert-proto/advert-proto"
)

type AdvertClient interface {
	GetAdverts(ctx context.Context, uuid string) (*advert.GetAdvertsOut, error)
	CreateAdvert(ctx context.Context, req *RequestData) (*advert.AdvertEmpty, error)
}
