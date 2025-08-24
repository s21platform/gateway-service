package materials

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/s21platform/materials-service/pkg/materials"

	"github.com/s21platform/gateway-service/internal/config"
)

type Service struct {
	client materials.MaterialsServiceClient
}

func New(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Materials.Host, cfg.Materials.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create grpc client: %v", err)
	}
	client := materials.NewMaterialsServiceClient(conn)
	return &Service{client: client}
}

func (s *Service) GetAllMaterials(ctx context.Context) (*materials.GetAllMaterialsOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))

	resp, err := s.client.GetAllMaterials(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get materials: %w", err)
	}

	return resp, nil
}
