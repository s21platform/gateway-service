package friends

import (
	"context"

	friends "github.com/s21platform/friends-proto/friends-proto"
)

type FriendsClient interface {
	GetCountFriends(ctx context.Context) (*friends.GetCountFriendsOut, error)
	SetFriends(ctx context.Context, peer *friends.SetFriendsIn) (*friends.SetFriendsOut, error)
	RemoveFriends(ctx context.Context, peer *friends.RemoveFriendsIn) (*friends.RemoveFriendsOut, error)
}
