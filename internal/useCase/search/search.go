package search

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/s21platform/search-proto/search"

	"github.com/s21platform/gateway-service/internal/model"
)

type UseCase struct {
	sS SearchClient
}

func New(sS SearchClient) *UseCase {
	return &UseCase{sS: sS}
}

func (u *UseCase) GetUsersWithLimit(r *http.Request) (model.SearchUsersOut, error) {
	var tmp model.SearchUsersOut
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	nickname := r.URL.Query().Get("nickname")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return tmp, fmt.Errorf("invalid limit: %v", err)
	}
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		return tmp, fmt.Errorf("invalid offset: %v", err)
	}
	resp, err := u.sS.GetUserWithLimit(r.Context(), &search.GetUserWithLimitIn{
		Limit:    limit,
		Offset:   offset,
		Nickname: nickname,
	})
	if err != nil {
		return tmp, fmt.Errorf("failed to call GetUserWithLimit(): %v", err)
	}
	tmp = model.SearchUsersOut{
		Users: make([]model.SearchUser, 0),
		Total: resp.Total,
	}
	for _, user := range resp.Users {
		users := model.SearchUser{
			Nickname:   user.Nickname,
			Uuid:       user.Uuid,
			AvatarLink: user.AvatarLink,
			Name:       user.Name,
			Surname:    user.Surname,
			IsFriend:   user.IsFriend,
		}
		tmp.Users = append(tmp.Users, users)
	}
	return tmp, nil
}

func (u *UseCase) GetSocietyWithLimit(r *http.Request) (model.SearchSocietyOut, error) {
	var tmp model.SearchSocietyOut
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	name := r.URL.Query().Get("name")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return tmp, fmt.Errorf("invalid limit: %v", err)
	}
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		return tmp, fmt.Errorf("invalid offset: %v", err)
	}
	res, err := u.sS.GetSocietyWithLimit(r.Context(), &search.GetSocietyWithLimitIn{
		Limit:  limit,
		Offset: offset,
		Name:   name,
	})
	if err != nil {
		return tmp, fmt.Errorf("failed to get society from search: %v", err)
	}
	tmp = model.SearchSocietyOut{
		Societies: make([]model.SearchSociety, 0),
		Total:     res.Total,
	}
	for _, societys := range res.Societies {
		society := model.SearchSociety{
			Name:       societys.Name,
			AvatarLink: societys.AvatarLink,
			SocietyId:  societys.SocietyId,
			IsMember:   societys.IsMember,
			IsPrivate:  societys.IsPrivate,
		}
		tmp.Societies = append(tmp.Societies, society)
	}
	return tmp, nil
}
