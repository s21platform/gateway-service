package notification

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	notificationproto "github.com/s21platform/notification-proto/notification-proto"

	"github.com/s21platform/gateway-service/internal/config"
)

type Client struct {
	client notificationproto.NotificationServiceClient
}

func New(cfg *config.Config) *Client {
	connStr := fmt.Sprintf("%s:%s", cfg.Notification.Host, cfg.Notification.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	client := notificationproto.NewNotificationServiceClient(conn)
	return &Client{client: client}
}

func (c *Client) GetCountNotification(ctx context.Context) (*notificationproto.NotificationCountOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	result, err := c.client.GetNotificationCount(ctx, &notificationproto.Empty{})
	if err != nil {
		log.Printf("failed to get notification count: %v", err)
		return nil, fmt.Errorf("failed to get notification count: %v", err)
	}
	return result, nil
}

func (c *Client) GetNotifications(ctx context.Context, limit int64, offset int64) (*notificationproto.NotificationOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	result, err := c.client.GetNotification(ctx, &notificationproto.NotificationIn{Limit: limit, Offset: offset})
	if err != nil {
		log.Printf("failed to get notifications: %v", err)
		return nil, fmt.Errorf("failed to get notifications: %v", err)
	}
	return result, nil
}
