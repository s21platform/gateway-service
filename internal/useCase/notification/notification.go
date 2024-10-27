package notification

import (
	"fmt"
	"net/http"
	"strconv"

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

func (u *Usecase) GetNotification(r *http.Request) (*notificationproto.NotificationOut, error) {
	query := r.URL.Query()
	limitStr := query.Get("limit")
	offsetStr := query.Get("offset")

	if limitStr == "" {
		limitStr = "10"
	}
	if offsetStr == "" {
		offsetStr = "0"
	}

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse limit: %w", err)
	}

	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse offset: %w", err)
	}

	result, err := u.nC.GetNotifications(r.Context(), limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get notifications: %w", err)
	}
	return result, nil
}
