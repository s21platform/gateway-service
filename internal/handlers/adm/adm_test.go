package adm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	logger_lib "github.com/s21platform/logger-lib"
	"github.com/s21platform/staff-service/pkg/staff"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/s21platform/gateway-service/internal/config"
)

func setupTestContext(r *http.Request, ctrl *gomock.Controller) *http.Request {
	mockLogger := logger_lib.NewMockLoggerInterface(ctrl)
	mockLogger.EXPECT().AddFuncName(gomock.Any()).AnyTimes()
	mockLogger.EXPECT().Error(gomock.Any()).AnyTimes()
	ctx := context.WithValue(r.Context(), config.KeyLogger, mockLogger)
	return r.WithContext(ctx)
}

func TestHandler_StaffLogin(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name           string
		requestBody    interface{}
		setupMock      func(mock *MockStaffService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful_login",
			requestBody: staff.LoginIn{
				Login:    "test@example.com",
				Password: "password123",
			},
			setupMock: func(mock *MockStaffService) {
				response := &staff.LoginOut{
					AccessToken:  "test-access-token",
					RefreshToken: "test-refresh-token",
					ExpiresAt:    1743536404,
					Staff: &staff.Staff{
						Id:        "test-id",
						Login:     "test@example.com",
						RoleId:    1,
						RoleName:  "admin",
						CreatedAt: 1743524411,
						UpdatedAt: 1743524411,
					},
				}
				mock.EXPECT().
					StaffLogin(gomock.Any()).
					Return(response, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{"access_token":"test-access-token","refresh_token":"test-refresh-token",` +
				`"expires_at":1743536404,"staff":{"id":"test-id","login":"test@example.com",` +
				`"role_id":1,"role_name":"admin","created_at":1743524411,"updated_at":1743524411}}`,
		},
		{
			name:        "invalid_request_body",
			requestBody: "invalid json",
			setupMock: func(mock *MockStaffService) {
				mock.EXPECT().
					StaffLogin(gomock.Any()).
					Return(nil, errors.New("failed to unmarshal request body"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "",
		},
		{
			name: "login_service_error",
			requestBody: staff.LoginIn{
				Login:    "test@example.com",
				Password: "password123",
			},
			setupMock: func(mock *MockStaffService) {
				mock.EXPECT().
					StaffLogin(gomock.Any()).
					Return(nil, errors.New("service error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockService := NewMockStaffService(ctrl)
			tt.setupMock(mockService)

			handler := New(mockService)

			var body []byte
			var err error
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, err = json.Marshal(tt.requestBody)
				require.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/adm/auth/login", bytes.NewReader(body))
			req = setupTestContext(req, ctrl)
			rec := httptest.NewRecorder()

			handler.StaffLogin(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedBody != "" {
				assert.JSONEq(t, tt.expectedBody, rec.Body.String())
			}
		})
	}
}

func TestHandler_CreateStaff(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name           string
		requestBody    interface{}
		setupMock      func(mock *MockStaffService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful_create",
			requestBody: staff.CreateIn{
				Login:    "new@example.com",
				Password: "newpassword123",
				RoleId:   1,
			},
			setupMock: func(mock *MockStaffService) {
				response := &staff.Staff{
					Id:        "new-id",
					Login:     "new@example.com",
					RoleId:    1,
					RoleName:  "admin",
					CreatedAt: 1743524411,
					UpdatedAt: 1743524411,
				}
				mock.EXPECT().
					CreateStaff(gomock.Any()).
					Return(response, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{"id":"new-id","login":"new@example.com",` +
				`"role_id":1,"role_name":"admin","created_at":1743524411,"updated_at":1743524411}`,
		},
		{
			name:        "invalid_request_body",
			requestBody: "invalid json",
			setupMock: func(mock *MockStaffService) {
				mock.EXPECT().
					CreateStaff(gomock.Any()).
					Return(nil, errors.New("failed to unmarshal request body"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "",
		},
		{
			name: "create_staff_error",
			requestBody: staff.CreateIn{
				Login:    "new@example.com",
				Password: "newpassword123",
				RoleId:   1,
			},
			setupMock: func(mock *MockStaffService) {
				mock.EXPECT().
					CreateStaff(gomock.Any()).
					Return(nil, errors.New("service error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockService := NewMockStaffService(ctrl)
			tt.setupMock(mockService)

			handler := New(mockService)

			var body []byte
			var err error
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, err = json.Marshal(tt.requestBody)
				require.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/adm/staff", bytes.NewReader(body))
			req = setupTestContext(req, ctrl)
			rec := httptest.NewRecorder()

			handler.CreateStaff(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedBody != "" {
				assert.JSONEq(t, tt.expectedBody, rec.Body.String())
			}
		})
	}
}

func TestHandler_ListStaff(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name           string
		setupRequest   func(*http.Request)
		setupMock      func(mock *MockStaffService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful_list",
			setupRequest: func(r *http.Request) {
				q := r.URL.Query()
				q.Set("page", "2")
				q.Set("page_size", "20")
				q.Set("search_term", "test")
				q.Set("role_id", "1")
				r.URL.RawQuery = q.Encode()
			},
			setupMock: func(mock *MockStaffService) {
				response := &staff.ListOut{
					Staff: []*staff.Staff{
						{
							Id:        "staff-1",
							Login:     "staff1@example.com",
							RoleId:    1,
							RoleName:  "admin",
							CreatedAt: 1743524411,
							UpdatedAt: 1743524411,
						},
					},
				}
				mock.EXPECT().
					ListStaff(gomock.Any()).
					Return(response, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{"staff":[{"id":"staff-1","login":"staff1@example.com",` +
				`"role_id":1,"role_name":"admin","created_at":1743524411,"updated_at":1743524411}]}`,
		},
		{
			name: "list_staff_error",
			setupRequest: func(r *http.Request) {
				// Не добавляем query-параметры
			},
			setupMock: func(mock *MockStaffService) {
				mock.EXPECT().
					ListStaff(gomock.Any()).
					Return(nil, errors.New("service error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockService := NewMockStaffService(ctrl)
			tt.setupMock(mockService)

			handler := New(mockService)

			req := httptest.NewRequest(http.MethodGet, "/adm/staff/list", nil)
			tt.setupRequest(req)
			req = setupTestContext(req, ctrl)
			rec := httptest.NewRecorder()

			handler.ListStaff(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedBody != "" {
				assert.JSONEq(t, tt.expectedBody, rec.Body.String())
			}
		})
	}
}

func TestHandler_GetStaff(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name           string
		staffID        string
		setupMock      func(mock *MockStaffService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:    "successful_get",
			staffID: "test-uuid",
			setupMock: func(mock *MockStaffService) {
				response := &staff.Staff{
					Id:        "test-uuid",
					Login:     "test@example.com",
					RoleId:    1,
					RoleName:  "admin",
					CreatedAt: 1743524411,
					UpdatedAt: 1743524411,
				}
				mock.EXPECT().
					GetStaff(gomock.Any(), "test-uuid").
					Return(response, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{"id":"test-uuid","login":"test@example.com",` +
				`"role_id":1,"role_name":"admin","created_at":1743524411,"updated_at":1743524411}`,
		},
		{
			name:    "empty_staff_id",
			staffID: "",
			setupMock: func(mock *MockStaffService) {
				mock.EXPECT().
					GetStaff(gomock.Any(), "").
					Return(nil, errors.New("staff ID is required"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "",
		},
		{
			name:    "service_error",
			staffID: "test-uuid",
			setupMock: func(mock *MockStaffService) {
				mock.EXPECT().
					GetStaff(gomock.Any(), "test-uuid").
					Return(nil, errors.New("service error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockService := NewMockStaffService(ctrl)
			tt.setupMock(mockService)

			handler := New(mockService)

			req := httptest.NewRequest(http.MethodGet, "/staff/"+tt.staffID, nil)
			req = setupTestContext(req, ctrl)
			rec := httptest.NewRecorder()

			if tt.staffID == "" {
				handler.GetStaff(rec, req)
			} else {
				r := chi.NewRouter()
				r.Get("/staff/{uuid}", handler.GetStaff)
				r.ServeHTTP(rec, req)
			}

			assert.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedBody != "" {
				assert.JSONEq(t, tt.expectedBody, rec.Body.String())
			}
		})
	}
}

func TestAttachAdmRoutes(t *testing.T) {
	t.Parallel()

	routes := []struct {
		method         string
		path           string
		expectedStatus int
	}{
		{http.MethodPost, "/adm/auth/login", http.StatusInternalServerError},
		{http.MethodPost, "/adm/staff", http.StatusUnauthorized},
		{http.MethodGet, "/adm/staff/list", http.StatusUnauthorized},
		{http.MethodGet, "/adm/staff/test-uuid", http.StatusUnauthorized},
	}

	for _, route := range routes {
		route := route
		t.Run(route.path, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mockService := NewMockStaffService(ctrl)

			// Настраиваем ожидания для мока в зависимости от маршрута
			if route.path == "/adm/auth/login" {
				mockService.EXPECT().
					StaffLogin(gomock.Any()).
					Return(nil, errors.New("invalid request")).
					AnyTimes()
			}

			handler := New(mockService)
			router := chi.NewRouter()
			AttachAdmRoutes(router, handler)

			req := httptest.NewRequest(route.method, route.path, nil)
			req = setupTestContext(req, ctrl)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			assert.Equal(t, route.expectedStatus, rec.Code)
		})
	}
}
