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
	"github.com/s21platform/staff-service/pkg/staff"
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
			requestBody: staff.LoginIn{
				Login:    "test@example.com",
				Password: "password123",
			},
			setupMock: func(mock *MockStaffClient) {
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
			requestBody: staff.LoginIn{
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

func TestHandler_CreateStaff(t *testing.T) {
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
			name: "successful_create",
			requestBody: staff.CreateIn{
				Login:    "new@example.com",
				Password: "newpassword123",
				RoleId:   1,
			},
			setupMock: func(mock *MockStaffClient) {
				response := &staff.Staff{
					Id:        "new-id",
					Login:     "new@example.com",
					RoleId:    1,
					RoleName:  "admin",
					CreatedAt: 1743524411,
					UpdatedAt: 1743524411,
				}
				mock.EXPECT().
					CreateStaff(gomock.Any(), gomock.Any()).
					Return(response, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{"id":"new-id","login":"new@example.com",` +
				`"role_id":1,"role_name":"admin","created_at":1743524411,"updated_at":1743524411}`,
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
			name: "create_staff_error",
			requestBody: staff.CreateIn{
				Login:    "new@example.com",
				Password: "newpassword123",
				RoleId:   1,
			},
			setupMock: func(mock *MockStaffClient) {
				mock.EXPECT().
					CreateStaff(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("service error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "failed to create staff\n",
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

			req := httptest.NewRequest(http.MethodPost, "/adm/staff", bytes.NewReader(body))
			rec := httptest.NewRecorder()

			handler.CreateStaff(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedStatus == http.StatusOK {
				assert.JSONEq(t, tt.expectedBody, rec.Body.String())
			} else {
				assert.Equal(t, tt.expectedBody, rec.Body.String())
			}
		})
	}
}

func TestHandler_ListStaff(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	searchTerm := "test"
	roleID := int32(1)

	tests := []struct {
		name           string
		setupRequest   func(*http.Request)
		setupMock      func(mock *MockStaffClient)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful_list_without_params",
			setupRequest: func(r *http.Request) {
				// Не добавляем query-параметры
			},
			setupMock: func(mock *MockStaffClient) {
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
					ListStaff(gomock.Any(), &staff.ListIn{
						Page:     1,
						PageSize: 10,
					}).
					Return(response, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{"staff":[{"id":"staff-1","login":"staff1@example.com",` +
				`"role_id":1,"role_name":"admin","created_at":1743524411,"updated_at":1743524411}]}`,
		},
		{
			name: "successful_list_with_all_params",
			setupRequest: func(r *http.Request) {
				q := r.URL.Query()
				q.Set("page", "2")
				q.Set("page_size", "20")
				q.Set("search_term", searchTerm)
				q.Set("role_id", "1")
				r.URL.RawQuery = q.Encode()
			},
			setupMock: func(mock *MockStaffClient) {
				response := &staff.ListOut{
					Staff: []*staff.Staff{
						{
							Id:        "staff-2",
							Login:     "staff2@example.com",
							RoleId:    1,
							RoleName:  "admin",
							CreatedAt: 1743524412,
							UpdatedAt: 1743524412,
						},
					},
				}
				mock.EXPECT().
					ListStaff(gomock.Any(), &staff.ListIn{
						Page:       2,
						PageSize:   20,
						SearchTerm: &searchTerm,
						RoleId:     &roleID,
					}).
					Return(response, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{"staff":[{"id":"staff-2","login":"staff2@example.com",` +
				`"role_id":1,"role_name":"admin","created_at":1743524412,"updated_at":1743524412}]}`,
		},
		{
			name: "invalid_page_parameter",
			setupRequest: func(r *http.Request) {
				q := r.URL.Query()
				q.Set("page", "invalid")
				r.URL.RawQuery = q.Encode()
			},
			setupMock: func(mock *MockStaffClient) {
				// Мок не нужен, так как ошибка произойдет при парсинге параметра
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid page parameter\n",
		},
		{
			name: "invalid_page_size_parameter",
			setupRequest: func(r *http.Request) {
				q := r.URL.Query()
				q.Set("page_size", "invalid")
				r.URL.RawQuery = q.Encode()
			},
			setupMock: func(mock *MockStaffClient) {
				// Мок не нужен, так как ошибка произойдет при парсинге параметра
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid page_size parameter\n",
		},
		{
			name: "invalid_role_id_parameter",
			setupRequest: func(r *http.Request) {
				q := r.URL.Query()
				q.Set("role_id", "invalid")
				r.URL.RawQuery = q.Encode()
			},
			setupMock: func(mock *MockStaffClient) {
				// Мок не нужен, так как ошибка произойдет при парсинге параметра
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid role_id parameter\n",
		},
		{
			name: "list_staff_error",
			setupRequest: func(r *http.Request) {
				// Не добавляем query-параметры
			},
			setupMock: func(mock *MockStaffClient) {
				mock.EXPECT().
					ListStaff(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("service error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "failed to list staff\n",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockClient := NewMockStaffClient(ctrl)
			tt.setupMock(mockClient)

			handler := New(mockClient)

			req := httptest.NewRequest(http.MethodGet, "/adm/staff/list", nil)
			tt.setupRequest(req)
			rec := httptest.NewRecorder()

			handler.ListStaff(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedStatus == http.StatusOK {
				assert.JSONEq(t, tt.expectedBody, rec.Body.String())
			} else {
				assert.Equal(t, tt.expectedBody, rec.Body.String())
			}
		})
	}
}

func TestAttachAdmRoutes(t *testing.T) {
	t.Parallel()

	routes := []struct {
		method string
		path   string
	}{
		{http.MethodPost, "/adm/auth/login"},
		{http.MethodPost, "/adm/staff"},
		{http.MethodGet, "/adm/staff/list"},
	}

	for _, route := range routes {
		route := route
		t.Run(route.path, func(t *testing.T) {
			t.Parallel()

			handler := New(NewMockStaffClient(gomock.NewController(t)))
			router := chi.NewRouter()
			AttachAdmRoutes(router, handler)

			req := httptest.NewRequest(route.method, route.path, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			// Проверяем, что маршрут существует (не возвращает 404)
			// Для /adm/auth/login ожидаем 400 Bad Request
			// Для /adm/staff и /adm/staff/list ожидаем 401 Unauthorized из-за middleware CheckJWT
			expectedStatus := http.StatusBadRequest
			if route.path == "/adm/staff" || route.path == "/adm/staff/list" {
				expectedStatus = http.StatusUnauthorized
			}
			assert.Equal(t, expectedStatus, rec.Code)
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
		setupMock      func(mock *MockStaffClient)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:    "successful_get",
			staffID: "test-uuid",
			setupMock: func(mock *MockStaffClient) {
				response := &staff.Staff{
					Id:        "test-uuid",
					Login:     "test@example.com",
					RoleId:    1,
					RoleName:  "admin",
					CreatedAt: 1743524411,
					UpdatedAt: 1743524411,
				}
				mock.EXPECT().
					GetStaff(gomock.Any(), &staff.GetIn{Id: "test-uuid"}).
					Return(response, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{"id":"test-uuid","login":"test@example.com",` +
				`"role_id":1,"role_name":"admin","created_at":1743524411,"updated_at":1743524411}` + "\n",
		},
		{
			name:    "empty_staff_id",
			staffID: "",
			setupMock: func(mock *MockStaffClient) {
				// Мок не нужен, так как ошибка произойдет при проверке ID
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "staff ID is required\n",
		},
		{
			name:    "service_error",
			staffID: "test-uuid",
			setupMock: func(mock *MockStaffClient) {
				mock.EXPECT().
					GetStaff(gomock.Any(), &staff.GetIn{Id: "test-uuid"}).
					Return(nil, errors.New("service error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "failed to get staff\n",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockClient := NewMockStaffClient(ctrl)
			tt.setupMock(mockClient)

			handler := New(mockClient)

			req := httptest.NewRequest(http.MethodGet, "/staff/"+tt.staffID, nil)
			rec := httptest.NewRecorder()

			// Для теста с пустым ID используем прямой вызов обработчика
			if tt.staffID == "" {
				handler.GetStaff(rec, req)
			} else {
				// Для остальных тестов используем роутер
				r := chi.NewRouter()
				r.Get("/staff/{uuid}", handler.GetStaff)
				r.ServeHTTP(rec, req)
			}

			assert.Equal(t, tt.expectedStatus, rec.Code)
			assert.Equal(t, tt.expectedBody, rec.Body.String())
		})
	}
}
