package chat

import (
	"context"
<<<<<<< Updated upstream

=======
>>>>>>> Stashed changes
	chat "github.com/s21platform/chat-proto/chat-proto"
)

type ChatClient interface {
	GetRecentMessages(ctx context.Context, uuid string) (*chat.GetRecentMessagesOut, error)
}
