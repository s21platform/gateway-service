package search

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/s21platform/search-proto/search"
)

type UseCase struct {
	sS SearchClient
}

func New(sS SearchClient) *UseCase {
	return &UseCase{sS: sS}
}

func (u *UseCase) GetUsersWithLimit(r *http.Request) (*search.GetUserWithLimitOut, error) {
	var readBody struct {
		Limit    int64  `json:"limit"`
		Offset   int64  `json:"offset"`
		Nickname string `json:"nickname"`
	}
	readType := r.URL.Query().Get("type")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if len(body) == 0 {
		return nil, fmt.Errorf("failed to request body is empty")
	}
	if err := json.Unmarshal(body, &readBody); err != nil {
		return nil, fmt.Errorf("failed to json.Unmarshal: %v", err)
	}
	resp, err := u.sS.GetUserWithLimit(r.Context(), &search.GetUserWithLimitIn{
		Limit:    readBody.Limit,
		Offset:   readBody.Offset,
		Nickname: readBody.Nickname,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to call GetUserWithLimit(): %v", err)
	}
	return resp, nil
}
