package notification

import (
	"context"

	notificationproto "github.com/s21platform/notification-proto/notification-proto"
)

type NotificationClient interface {
	GetNotifications(ctx context.Context, limit int64, offset int64)
	GetCountNotification(ctx context.Context) (*notificationproto.NotificationCountOut, error)
}
