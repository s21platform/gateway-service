package advert

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	advert "github.com/s21platform/advert-proto/advert-proto"
	model "github.com/s21platform/gateway-service/internal/model"
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

func (u *Usecase) CreateAdvert(r *http.Request) (*advert.AdvertEmpty, error) {
	requestData := model.CreateAdvertRequestData{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %v", err)
	}
	defer r.Body.Close()

	if len(body) == 0 {
		return nil, fmt.Errorf("request body is empty")
	}

	err = json.Unmarshal(body, &requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode request body: %v", err)
	}

	resp, err := u.aC.CreateAdvert(r.Context(), &requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to create advert in usecase: %v", err)
	}
	return resp, nil
}

func (u *Usecase) CancelAdvert(r *http.Request) (*advert.AdvertEmpty, error) {
	requestData := model.CancelAdvertRequestData{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %v", err)
	}
	defer r.Body.Close()

	if len(body) == 0 {
		return nil, fmt.Errorf("request body is empty")
	}

	err = json.Unmarshal(body, &requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode request body: %v", err)
	}

	resp, err := u.aC.CancelAdvert(r.Context(), &requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to cancel advert in usecase: %v", err)
	}
	return resp, nil
}
