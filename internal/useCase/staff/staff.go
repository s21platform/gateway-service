package staff

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	logger_lib "github.com/s21platform/logger-lib"
	"github.com/s21platform/staff-service/pkg/staff"

	"github.com/s21platform/gateway-service/internal/config"
)

type UseCase struct {
	client StaffClient
}

func New(client StaffClient) *UseCase {
	return &UseCase{client: client}
}

func (u *UseCase) StaffLogin(r *http.Request) (*staff.LoginOut, error) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to read request body: %v", err))
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	var loginRequest staff.LoginIn
	if err := json.Unmarshal(body, &loginRequest); err != nil {
		logger.Error(fmt.Sprintf("failed to unmarshal request body: %v", err))
		return nil, fmt.Errorf("failed to unmarshal request body: %w", err)
	}

	return u.client.StaffLogin(r.Context(), &loginRequest)
}

func (u *UseCase) CreateStaff(r *http.Request) (*staff.Staff, error) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to read request body: %v", err))
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	var req staff.CreateIn
	if err := json.Unmarshal(body, &req); err != nil {
		logger.Error(fmt.Sprintf("failed to unmarshal request body: %v", err))
		return nil, fmt.Errorf("failed to unmarshal request body: %w", err)
	}

	return u.client.CreateStaff(r.Context(), &req)
}

func (u *UseCase) ListStaff(r *http.Request) (*staff.ListOut, error) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)

	req := &staff.ListIn{
		Page:     1,
		PageSize: 10,
	}

	if page := r.URL.Query().Get("page"); page != "" {
		pageInt, err := strconv.ParseInt(page, 10, 32)
		if err != nil {
			logger.Error(fmt.Sprintf("invalid page parameter: %v", err))
			return nil, fmt.Errorf("invalid page parameter: %w", err)
		}
		req.Page = int32(pageInt)
	}

	if pageSize := r.URL.Query().Get("page_size"); pageSize != "" {
		pageSizeInt, err := strconv.ParseInt(pageSize, 10, 32)
		if err != nil {
			logger.Error(fmt.Sprintf("invalid page_size parameter: %v", err))
			return nil, fmt.Errorf("invalid page_size parameter: %w", err)
		}
		req.PageSize = int32(pageSizeInt)
	}

	if searchTerm := r.URL.Query().Get("search_term"); searchTerm != "" {
		req.SearchTerm = &searchTerm
	}

	if roleID := r.URL.Query().Get("role_id"); roleID != "" {
		roleIDInt, err := strconv.ParseInt(roleID, 10, 32)
		if err != nil {
			logger.Error(fmt.Sprintf("invalid role_id parameter: %v", err))
			return nil, fmt.Errorf("invalid role_id parameter: %w", err)
		}
		roleIDInt32 := int32(roleIDInt)
		req.RoleId = &roleIDInt32
	}

	return u.client.ListStaff(r.Context(), req)
}

func (u *UseCase) GetStaff(r *http.Request, staffID string) (*staff.Staff, error) {
	if staffID == "" {
		return nil, fmt.Errorf("staff ID is required")
	}

	return u.client.GetStaff(r.Context(), &staff.GetIn{Id: staffID})
}
