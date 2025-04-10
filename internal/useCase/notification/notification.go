package notification

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	logger_lib "github.com/s21platform/logger-lib"
	"github.com/s21platform/notification-service/pkg/notification"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/model"
)

type UseCase struct {
	client NotificationClient
}

func New(client NotificationClient) *UseCase {
	return &UseCase{client: client}
}

func (u *UseCase) GetCountNotification(r *http.Request) (*notification.NotificationCountOut, error) {
	return u.client.GetCountNotification(r.Context())
}

func (u *UseCase) GetNotification(r *http.Request) (*notification.NotificationOut, error) {
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

	return u.client.GetNotifications(r.Context(), limit, offset)
}

func (u *UseCase) MarkNotificationAsRead(r *http.Request) (*emptypb.Empty, error) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to read request body: %v", err))
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	var req model.MarkNotificationsAsReadRequest
	if err := json.Unmarshal(body, &req); err != nil {
		logger.Error(fmt.Sprintf("failed to unmarshal request body: %v", err))
		return nil, fmt.Errorf("failed to unmarshal request body: %w", err)
	}

	if len(req.Data.IDs) == 0 {
		logger.Error("notification IDs are required")
		return nil, fmt.Errorf("notification IDs are required")
	}

	return u.client.MarkNotificationsAsRead(r.Context(), req.Data.IDs)
}
