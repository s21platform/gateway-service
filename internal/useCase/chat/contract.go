package chat

import (
	"context"

	"github.com/s21platform/chat-service/pkg/chat"
)

type ChatClient interface {
	GetChats(ctx context.Context) (*chat.GetChatsOut, error)
	CreatePrivateChat(ctx context.Context, uuid string) (*chat.CreatePrivateChatOut, error)
	GetPrivateRecentMessages(ctx context.Context, uuid string) (*chat.GetPrivateRecentMessagesOut, error)
}
