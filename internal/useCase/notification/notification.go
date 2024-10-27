package notification

import (
	"fmt"
	"net/http"

	notificationproto "github.com/s21platform/notification-proto/notification-proto"
)

type Usecase struct {
	nC NotificationClient
}

func New(nC NotificationClient) *Usecase {
	return &Usecase{nC: nC}
}

func (u *Usecase) GetCountNotification(r *http.Request) (*notificationproto.NotificationCountOut, error) {
	result, err := u.nC.GetCountNotification(r.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get notification count: %w", err)
	}
	return result, nil
}
