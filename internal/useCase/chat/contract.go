package chat

import (
	"context"

	chat "github.com/s21platform/chat-proto/chat-proto"
)

type ChatClient interface {
	CreatePrivateChat(ctx context.Context, uuid string) (*chat.CreatePrivateChatOut, error)
	//GetRecentMessages(ctx context.Context, uuid string) (*chat.GetRecentMessagesOut, error)
}
