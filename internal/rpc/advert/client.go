package advert

import (
	"context"
	"fmt"
	"github.com/s21platform/gateway-service/internal/useCase/advert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	advertproto "github.com/s21platform/advert-proto/advert-proto"

	"github.com/s21platform/gateway-service/internal/config"
)

type Service struct {
	client advertproto.AdvertServiceClient
}

func New(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Advert.Host, cfg.Advert.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create grpc client: %v", err)
	}
	client := advertproto.NewAdvertServiceClient(conn)
	return &Service{client: client}
}

func (s *Service) GetAdverts(ctx context.Context, uuid string) (*advertproto.GetAdvertsOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", uuid))

	resp, err := s.client.GetAdverts(ctx, &advertproto.AdvertEmpty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get adverts in grpc: %w", err)
	}

	return resp, nil
}

func (s *Service) CreateAdvrt(ctx context.Context, req *advert.RequestData) (*advertproto.AdvertEmpty, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	request := &advertproto.CreateAdvertIn{
		Text: req.TextContent,
		User: &advertproto.UserFilter{
			Os: req.UserFilter.Os,
		},
		ExpiredAt: timestamppb.New(req.ExpiredAt),
	}

	resp, err := s.client.CreateAdvert(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to create advert in grpc: %w", err)
	}

	return resp, nil
}
