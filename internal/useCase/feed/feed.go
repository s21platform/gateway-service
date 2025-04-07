package feed

import (
	"encoding/json"
	"fmt"
	"net/http"

	feed "github.com/s21platform/feed-proto/feed-proto"
	"github.com/s21platform/gateway-service/internal/model"
)

type Usecase struct {
	feS FeedClient
}

func New(feS FeedClient) *Usecase {
	return &Usecase{feS: feS}
}

func (u *Usecase) CreateUserPost(r *http.Request) (*feed.CreateUserPostOut, error) {
	requestData := model.CreateUserPostRequestData{}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode request body: %v", err)
	}
	defer r.Body.Close()

	resp, err := u.feS.CreateUserPost(r.Context(), &requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to create post in usecase: %v", err)
	}
	return resp, nil
}
