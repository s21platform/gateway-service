package adm

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	api "github.com/s21platform/staff-service/pkg/staff/v0"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_StaffLogin(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name           string
		requestBody    interface{}
		setupMock      func(mock *MockStaffClient)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful_login",
			requestBody: api.LoginRequest{
				Login:    "test@example.com",
				Password: "password123",
			},
			setupMock: func(mock *MockStaffClient) {
				response := &api.LoginResponse{
					AccessToken:  "test-access-token",
					RefreshToken: "test-refresh-token",
					ExpiresAt:    1743536404,
					Staff: &api.Staff{
						Id:        "test-id",
						Login:     "test@example.com",
						RoleId:    1,
						RoleName:  "admin",
						CreatedAt: 1743524411,
						UpdatedAt: 1743524411,
					},
				}
				mock.EXPECT().
					StaffLogin(gomock.Any(), gomock.Any()).
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
			setupMock: func(mock *MockStaffClient) {
				// Мок не нужен, так как ошибка произойдет при декодировании
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "failed to decode request body\n",
		},
		{
			name: "login_service_error",
			requestBody: api.LoginRequest{
				Login:    "test@example.com",
				Password: "password123",
			},
			setupMock: func(mock *MockStaffClient) {
				mock.EXPECT().
					StaffLogin(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("service error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "failed to login\n",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockClient := NewMockStaffClient(ctrl)
			tt.setupMock(mockClient)

			handler := New(mockClient)

			var body []byte
			var err error
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, err = json.Marshal(tt.requestBody)
				require.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/adm/auth/login", bytes.NewReader(body))
			rec := httptest.NewRecorder()

			handler.StaffLogin(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			assert.Equal(t, tt.expectedBody, rec.Body.String())
		})
	}
}

func TestCheckJWT(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		setupRequest   func(*http.Request)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "no_auth_header",
			setupRequest: func(r *http.Request) {
				// Не добавляем заголовок Authorization
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Unauthorized\n",
		},
		{
			name: "invalid_auth_header_format",
			setupRequest: func(r *http.Request) {
				r.Header.Set("Authorization", "invalid-format")
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Unauthorized\n",
		},
		{
			name: "invalid_bearer_format",
			setupRequest: func(r *http.Request) {
				r.Header.Set("Authorization", "Bearer")
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Unauthorized\n",
		},
		{
			name: "valid_token",
			setupRequest: func(r *http.Request) {
				r.Header.Set("Authorization", "Bearer valid-token")
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name: "skip_auth_path",
			setupRequest: func(r *http.Request) {
				r.URL.Path = "/adm/auth/login"
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			tt.setupRequest(req)
			rec := httptest.NewRecorder()

			middleware := CheckJWT(nextHandler)
			middleware.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			assert.Equal(t, tt.expectedBody, rec.Body.String())
		})
	}
}

func TestAttachAdmRoutes(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockStaffClient(ctrl)
	handler := New(mockClient)

	r := chi.NewRouter()
	AttachAdmRoutes(r, handler)

	// Проверяем, что маршрут /adm/auth/login существует и использует правильный метод
	routes := []struct {
		method string
		path   string
	}{
		{http.MethodPost, "/adm/auth/login"},
	}

	for _, route := range routes {
		route := route
		t.Run(route.path, func(t *testing.T) {
			t.Parallel()

			// Создаем тестовый запрос
			req := httptest.NewRequest(route.method, route.path, nil)
			rec := httptest.NewRecorder()

			// Выполняем запрос через роутер
			r.ServeHTTP(rec, req)

			// Проверяем, что маршрут существует (не возвращает 404)
			// Мы ожидаем 400 Bad Request из-за middleware CheckJWT
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		})
	}
}
