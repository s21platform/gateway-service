package chat

import (
	"context"
	"fmt"
	"log"

	chat_proto "github.com/s21platform/chat-proto/chat-proto"
	"github.com/s21platform/gateway-service/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type Service struct {
	client chat_proto.ChatServiceClient
}

func NewService(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Chat.Host, cfg.Chat.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create gRPC client: %v", err)
	}
	client := chat_proto.NewChatServiceClient(conn)
	return &Service{client: client}
}

func (s *Service) GetRecentMessages(ctx context.Context, uuid string) (*chat_proto.GetRecentMessagesOut, error) {
	req := chat_proto.GetRecentMessagesIn{
		Uuid: uuid,
	}
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	resp, err := s.client.GetRecentMessages(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to GetRecentMessages in rpc: %v", err)
	}
	return resp, nil
}
