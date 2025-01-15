package chat

import (
	"context"
	
	chat "github.com/s21platform/chat-proto/chat-proto"
)

type ChatClient interface {
	GetRecentMessages(ctx context.Context, in *chat.GetRecentMessagesIn) (*chat.GetRecentMessagesOut, error)
}
