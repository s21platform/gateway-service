package feed

import (
	"fmt"
	"net/http"

	feed "github.com/s21platform/feed-proto/feed-proto"
)

type Usecase struct {
	feS FeedClient
}

func New(feS FeedClient) *Usecase {
	return &Usecase{feS: feS}
}

func (u *Usecase) CreateUserPost(r *http.Request) (*feed.CreateUserPostOut, error) {
	content := r.URL.Query().Get("content")

	resp, err := u.feS.CreateUserPost(r.Context(), content)
	if err != nil {
		return nil, fmt.Errorf("failed to create post in usecase: %v", err)
	}
	return resp, nil
}
