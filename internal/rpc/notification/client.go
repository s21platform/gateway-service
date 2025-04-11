package notification

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/s21platform/notification-service/pkg/notification"

	"github.com/s21platform/gateway-service/internal/config"
)

type Client struct {
	client notification.NotificationServiceClient
}

func New(cfg *config.Config) *Client {
	connStr := fmt.Sprintf("%s:%s", cfg.Notification.Host, cfg.Notification.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	client := notification.NewNotificationServiceClient(conn)
	return &Client{client: client}
}

func (c *Client) GetCountNotification(ctx context.Context) (*notification.NotificationCountOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	result, err := c.client.GetNotificationCount(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("failed to get notification count: %v", err)
		return nil, fmt.Errorf("failed to get notification count: %v", err)
	}
	return result, nil
}

func (c *Client) GetNotifications(ctx context.Context, limit int64, offset int64) (*notification.NotificationOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))
	result, err := c.client.GetNotification(ctx, &notification.NotificationIn{Limit: limit, Offset: offset})
	if err != nil {
		log.Printf("failed to get notifications: %v", err)
		return nil, fmt.Errorf("failed to get notifications: %w", err)
	}
	return result, nil
}

func (c *Client) MarkNotificationsAsRead(ctx context.Context, ids []int64) (*emptypb.Empty, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", ctx.Value(config.KeyUUID).(string)))

	_, err := c.client.MarkNotificationsAsRead(ctx, &notification.MarkNotificationsAsReadIn{NotificationIds: ids})
	if err != nil {
		return nil, fmt.Errorf("failed to mark notifications as read: %w", err)
	}

	return &emptypb.Empty{}, nil
}
