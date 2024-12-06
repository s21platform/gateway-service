package friends

import (
	"encoding/json"
	"fmt"
	"io"
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
	resp, err := u.fC.GetCountFriends(r.Context())
	if err != nil {
		return nil, fmt.Errorf("u.fC.GetCountFriends: %v", err)
	}
	return resp, nil
}

func (u *Usecase) SetFriends(r *http.Request) (*friends.SetFriendsOut, error) {
	var readPeer struct {
		Peer string `json:"peer"`
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("r.Body.ReadAll: %v", err)
	}
	defer r.Body.Close()
	if len(body) == 0 {
		return nil, fmt.Errorf("request body is empty")
	}
	if err = json.Unmarshal(body, &readPeer); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	resp, err := u.fC.SetFriends(r.Context(), &friends.SetFriendsIn{Peer: readPeer.Peer})
	if err != nil {
		return nil, fmt.Errorf("u.fC.SetFriends: %v", err)
	}
	return resp, nil
}

func (u *Usecase) RemoveFriends(r *http.Request) (*friends.RemoveFriendsOut, error) {
	return nil, nil
}
