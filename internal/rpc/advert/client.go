package advert

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	advert "github.com/s21platform/advert-proto/advert-proto"

	"github.com/s21platform/gateway-service/internal/config"
)

type Service struct {
	client advert.AdvertServiceClient
}

func New(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Advert.Host, cfg.Advert.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create grpc client: %v", err)
	}
	client := advert.NewAdvertServiceClient(conn)
	return &Service{client: client}
}

func (s *Service) GetAdverts(ctx context.Context, uuid string) (*advert.GetAdvertsOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", uuid))

	resp, err := s.client.GetAdverts(ctx, &advert.AdvertEmpty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get adverts in grpc: %w", err)
	}

	return resp, nil
}
