package society

import (
	"context"
	"fmt"
	"github.com/s21platform/gateway-service/internal/model"
	"log"

	"github.com/s21platform/gateway-service/internal/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	societyproto "github.com/s21platform/society-proto/society-proto"

	"github.com/s21platform/gateway-service/internal/config"
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

func (s *Service) CreateSociety(ctx context.Context, req *model.RequestData) (*societyproto.SetSocietyOut, error) {
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

func (s *Service) UpdateSociety(ctx context.Context, req *model.SocietyUpdate) error {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))

	var tags []*societyproto.TagsID
	if len(req.Tags) != 0 {
		for _, tag := range req.Tags {
			tags = append(tags, &societyproto.TagsID{TagID: tag})
		}
	}
	request := &societyproto.UpdateSocietyIn{
		SocietyUUID:    req.SocietyUUID,
		Name:           req.Name,
		Description:    req.Description,
		FormatID:       req.FormatID,
		PostPermission: req.PostPermissionID,
		IsSearch:       req.IsSearch,
		TagsID:         tags,
	}
	_, err := s.client.UpdateSociety(ctx, request)
	if err != nil {
		return fmt.Errorf("failed to update society: %v", err)
	}
	return nil
}
func (s *Service) RemoveSociety(ctx context.Context, req *model.SocietyId) (*societyproto.EmptySociety, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	request := &societyproto.RemoveSocietyIn{SocietyUUID: req.Id}
	resp, err := s.client.RemoveSociety(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to remove society: %v", err)
	}
	return resp, err
}

//func (s *Service) GetAccessLevel(ctx context.Context) (*societyproto.GetAccessLevelOut, error) {
//	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
//	resp, err := s.client.GetAccessLevel(ctx, &societyproto.EmptySociety{})
//	if err != nil {
//		return nil, fmt.Errorf("failed to get access level: %v", err)
//	}
//	return resp, nil
//}
//
//func (s *Service) GetSocietyInfo(ctx context.Context, id int64) (*societyproto.GetSocietyInfoOut, error) {
//	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
//	request := &societyproto.GetSocietyInfoIn{
//		Id: id,
//	}
//
//	resp, err := s.client.GetSocietyInfo(ctx, request)
//	if err != nil {
//		return nil, fmt.Errorf("failed to get society info: %v", err)
//	}
//	return resp, nil
//}
//
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
//func (s *Service) GetPermission(ctx context.Context) (*societyproto.GetPermissionsOut, error) {
//	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
//	resp, err := s.client.GetPermissions(ctx, &societyproto.EmptySociety{})
//	if err != nil {
//		return nil, fmt.Errorf("failed to get permissions: %v", err)
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
