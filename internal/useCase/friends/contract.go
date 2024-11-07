package friends

import (
	"context"

	friends "github.com/s21platform/friends-proto/friends-proto"
)

type FriendsClient interface {
	GetCountFriends(ctx context.Context, in *friends.Empty) (*friends.GetCountFriendsOut, error)
}
