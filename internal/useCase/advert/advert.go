package advert

import (
	"fmt"
	"net/http"

	advert "github.com/s21platform/advert-proto/advert-proto"
)

type Usecase struct {
	aC AdvertClient
}

func New(aC AdvertClient) *Usecase {
	return &Usecase{aC: aC}
}

func (u *Usecase) GetAdverts(r *http.Request) (*advert.GetAdvertsOut, error) {
	uuid := r.URL.Query().Get("uuid")

	resp, err := u.aC.GetAdverts(r.Context(), uuid)
	if err != nil {
		return nil, fmt.Errorf("failed get adverts in usecase: %v", err)
	}

	return resp, nil
}
