package notification

import (
	"context"

	"github.com/s21platform/notification-service/pkg/notification"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NotificationClient interface {
	GetCountNotification(ctx context.Context) (*notification.NotificationCountOut, error)
	GetNotifications(ctx context.Context, limit int64, offset int64) (*notification.NotificationOut, error)
	MarkNotificationsAsRead(ctx context.Context, ids []int64) (*emptypb.Empty, error)
}
