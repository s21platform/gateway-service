package society

import (
	"context"
	"fmt"
	"log"

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

func (s *Service) CreateSociety(ctx context.Context, name string, desc string, isPrivate bool, dirID int64, accessLevel int64) (*society_proto.SetSocietyOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	request := &society_proto.SetSocietyIn{
		Name:          name,
		Description:   desc,
		IsPrivate:     isPrivate,
		DirectionId:   dirID,
		AccessLevelId: accessLevel,
	}

	resp, err := s.client.CreateSociety(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("s.client.CreateSociety: %v", err)
	}
	log.Println("resp: ", resp)
	return resp, nil
}
