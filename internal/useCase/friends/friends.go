package friends

import (
	"fmt"
	"net/http"

	friends "github.com/s21platform/friends-proto/friends-proto"
)

type Usecase struct {
	fC FriendsClient
}

func New(fC FriendsClient) *Usecase {
	return &Usecase{fC: fC}
}

func (u *Usecase) GetCountFriends(r *http.Request) (*friends.GetCountFriendsOut, error) {
	req := &friends.Empty{}
	resp, err := u.fC.GetCountFriends(r.Context(), req)
	if err != nil {
		return nil, fmt.Errorf("u.fC.GetCountFriends: %v", err)
	}
	return resp, nil
}
