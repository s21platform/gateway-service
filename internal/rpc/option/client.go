package option

import (
	"context"
	"fmt"
	"log"

	"github.com/s21platform/gateway-service/internal/config"
	optionhub "github.com/s21platform/optionhub-proto/optionhub-proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func (s *Service) GetOSByID(ctx context.Context, id int64) (*optionhub.GetByIdOut, error) {
	req := optionhub.GetByIdIn{
		Id: id,
	}

	resp, err := s.client.GetOsById(ctx, &req) // эта функция наебнется
	if err != nil {
		return nil, fmt.Errorf("failed to get os in grpc: %w", err)
	}

	return resp, nil
}
