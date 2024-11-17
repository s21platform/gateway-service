package society

import (
	"context"
	"fmt"
	"log"

	"github.com/s21platform/gateway-service/internal/useCase/society"

	"google.golang.org/grpc/metadata"

	"github.com/s21platform/gateway-service/internal/config"
	society_proto "github.com/s21platform/society-proto/society-proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	client society_proto.SocietyServiceClient
}

func NewService(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Society.Host, cfg.Society.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create gRPC client: %v", err)
	}
	client := society_proto.NewSocietyServiceClient(conn)
	return &Service{client: client}
}

func (s *Service) CreateSociety(ctx context.Context, req *society.RequestData) (*society_proto.SetSocietyOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	request := &society_proto.SetSocietyIn{
		Name:          req.Name,
		Description:   req.Description,
		IsPrivate:     req.IsPrivate,
		DirectionId:   req.DirectionId,
		AccessLevelId: req.AccessLevelId,
	}

	resp, err := s.client.CreateSociety(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed grpc request: %v", err)
	}
	log.Println("resp: ", resp)
	return resp, nil
}
