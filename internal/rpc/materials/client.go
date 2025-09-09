package materials

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	materialsproto "github.com/s21platform/materials-service/pkg/materials"

	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/model"
)

type Service struct {
	client materialsproto.MaterialsServiceClient
}

func New(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Materials.Host, cfg.Materials.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create grpc client: %v", err)
	}
	client := materialsproto.NewMaterialsServiceClient(conn)
	return &Service{client: client}
}

func (s *Service) EditMaterial(ctx context.Context, req *model.EditMaterialRequest) (*materialsproto.EditMaterialOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))

	request := &materialsproto.EditMaterialIn{
		Uuid:            req.UUID,
		Title:           req.Title,
		CoverImageUrl:   req.CoverImageURL,
		Description:     req.Description,
		Content:         req.Content,
		ReadTimeMinutes: req.ReadTimeMinutes,
	}

	resp, err := s.client.EditMaterial(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to edit material in grpc: %w", err)
	}

	return resp, nil
}

func (s *Service) GetAllMaterials(ctx context.Context) (*materialsproto.GetAllMaterialsOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))

	resp, err := s.client.GetAllMaterials(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get materials: %w", err)
	}

	return resp, nil
}

func (s *Service) DeleteMaterial(ctx context.Context, uuid string) error {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))

	req := &materialsproto.DeleteMaterialIn{Uuid: uuid}

	_, err := s.client.DeleteMaterial(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to delete material: %w", err)
	}
	return nil
}
