package community

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	communityproto "github.com/s21platform/community-service/pkg/community"

	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/model"
)

type Service struct {
	client communityproto.CommunityServiceClient
}

func New(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Community.Host, cfg.Community.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create grpc client: %v", err)
	}
	client := communityproto.NewCommunityServiceClient(conn)
	return &Service{client: client}
}

func (s *Service) SendEduLinkingCode(ctx context.Context, in *model.SendEduLinkingCodeRequestData) (*emptypb.Empty, error) {
	request := &communityproto.SendEduLinkingCodeIn{
		Login: in.Login,
	}

	resp, err := s.client.SendEduLinkingCode(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to send edu linking code in grpc: %w", err)
	}

	return resp, nil
}
