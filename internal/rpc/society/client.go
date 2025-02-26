package society

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	societyproto "github.com/s21platform/society-proto/society-proto"

	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/useCase/society"
)

type Service struct {
	client societyproto.SocietyServiceClient
}

func NewService(cfg *config.Config) *Service {
	connStr := fmt.Sprintf("%s:%s", cfg.Society.Host, cfg.Society.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create gRPC client: %v", err)
	}
	client := societyproto.NewSocietyServiceClient(conn)
	return &Service{client: client}
}

func (s *Service) CreateSociety(ctx context.Context, req *society.RequestData) (*societyproto.SetSocietyOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	request := &societyproto.SetSocietyIn{
		Name:             req.Name,
		FormatID:         req.FormatID,
		PostPermissionID: req.PostPermissionID,
		IsSearch:         req.IsSearch,
	}

	resp, err := s.client.CreateSociety(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to create society: %v", err)
	}
	log.Println("resp: ", resp)
	return resp, nil
}

func (s *Service) GetSocietyInfo(ctx context.Context, societyUUID string) (*societyproto.GetSocietyInfoOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	req := &societyproto.GetSocietyInfoIn{SocietyUUID: societyUUID}

	resp, err := s.client.GetSocietyInfo(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get society info: %v", err)
	}
	return resp, nil
}

//func (s *Service) SubscribeToSociety(ctx context.Context, id int64) (*societyproto.SubscribeToSocietyOut, error) {
//	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
//	request := &societyproto.SubscribeToSocietyIn{
//		SocietyId: id,
//	}
//
//	resp, err := s.client.SubscribeToSociety(ctx, request)
//	if err != nil {
//		return nil, fmt.Errorf("failed subscribe to society: %v", err)
//	}
//	return resp, nil
//}
//
//func (s *Service) UnsubscribeFromSociety(ctx context.Context, id int64) (*societyproto.UnsubscribeFromSocietyOut, error) {
//	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
//	request := &societyproto.UnsubscribeFromSocietyIn{
//		SocietyId: id,
//	}
//
//	resp, err := s.client.UnsubscribeFromSociety(ctx, request)
//	if err != nil {
//		return nil, fmt.Errorf("failed to unsubscribe from society error: %v", err)
//	}
//	return resp, nil
//}
//
//func (s *Service) GetSocietiesForUser(ctx context.Context, uuid string) (*societyproto.GetSocietiesForUserOut, error) {
//	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
//
//	request := &societyproto.GetSocietiesForUserIn{
//		UserUuid: uuid,
//	}
//	resp, err := s.client.GetSocietiesForUser(ctx, request)
//	if err != nil {
//		return nil, fmt.Errorf("failed to get society for user error: %v", err)
//	}
//	return resp, err
//}
