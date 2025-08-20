package materials

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/s21platform/materials-service/pkg/materials"

	"github.com/s21platform/gateway-service/internal/config"
)

type Service struct {
	client materials.MaterialsServiceClient
}

func NewService(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Chat.Host, cfg.Chat.Port)

	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create gRPC client: %v", err)
	}
	client := materials.NewMaterialsServiceClient(conn)

	return &Service{client: client}
}

func (s *Service) EditMaterial(ctx context.Context, uuid string) (*materials.EditMaterialOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))

	req := materials.EditMaterialIn{
		Uuid: uuid,
	}

	resp, err := s.client.EditMaterial(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to edit material in rpc: %v", err)
	}

	return resp, nil
}
