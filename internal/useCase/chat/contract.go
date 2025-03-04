package chat

import (
	"context"

	chat "github.com/s21platform/chat-proto/chat-proto"
)

type ChatClient interface {
	CreatePrivateChat(ctx context.Context, uuid string) (*chat.CreatePrivateChatOut, error)
	GetPrivateRecentMessages(ctx context.Context, uuid string) (*chat.GetPrivateRecentMessagesOut, error)
}
