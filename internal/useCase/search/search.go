package search

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/s21platform/search-proto/search"
)

type UseCase struct {
	sS SearchClient
}

func New(sS SearchClient) *UseCase {
	return &UseCase{sS: sS}
}

func (u *UseCase) GetUsersWithLimit(r *http.Request) (*search.GetUserWithLimitOut, error) {
	readType := r.URL.Query().Get("type")
	if readType != "peer" {
		return nil, nil
	}
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	nickname := r.URL.Query().Get("nickname")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid limit: %v", err)
	}
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid offset: %v", err)
	}
	resp, err := u.sS.GetUserWithLimit(r.Context(), &search.GetUserWithLimitIn{
		Limit:    limit,
		Offset:   offset,
		Nickname: nickname,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to call GetUserWithLimit(): %v", err)
	}
	return resp, nil
}
