package chat

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/s21platform/chat-service/pkg/chat"

	"github.com/s21platform/gateway-service/internal/config"
)

type Service struct {
	client chat.ChatServiceClient
}

func NewService(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Chat.Host, cfg.Chat.Port)

	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create gRPC client: %v", err)
	}
	client := chat.NewChatServiceClient(conn)

	return &Service{client: client}
}

func (s *Service) GetChats(ctx context.Context) (*chat.GetChatsOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))

	resp, err := s.client.GetChats(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get chats in rpc: %v", err)
	}

	return resp, nil
}

func (s *Service) CreatePrivateChat(ctx context.Context, uuid string) (*chat.CreatePrivateChatOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))

	req := chat.CreatePrivateChatIn{
		CompanionUuid: uuid,
	}

	resp, err := s.client.CreatePrivateChat(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to create private chat in rpc: %v", err)
	}

	return resp, nil
}

func (s *Service) GetPrivateRecentMessages(ctx context.Context, uuid string) (*chat.GetPrivateRecentMessagesOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))

	req := chat.GetPrivateRecentMessagesIn{
		ChatUuid: uuid,
	}

	resp, err := s.client.GetPrivateRecentMessages(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to get private recent messages in rpc: %v", err)
	}

	return resp, nil
}
