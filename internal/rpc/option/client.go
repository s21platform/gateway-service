package option

import (
	"context"
	"fmt"
	"log"

	"github.com/s21platform/gateway-service/internal/config"
	optionhub "github.com/s21platform/optionhub-proto/optionhub-proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type Service struct {
	client optionhub.OptionhubServiceClient
}

func New(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Option.Host, cfg.Option.Port)

	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create grpc client: %v", err)
	}

	client := optionhub.NewOptionhubServiceClient(conn)

	return &Service{client: client}
}

func (s *Service) GetOsBySearchName(ctx context.Context, searchName *optionhub.GetByNameIn) (*optionhub.GetByNameOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))

	resp, err := s.client.GetOsBySearchName(ctx, searchName)
	if err != nil {
		return nil, fmt.Errorf("failed to get os list in grpc: %w", err)
	}

	return resp, nil
}
