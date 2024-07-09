package RESTHandlers

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	auth_proto "github.com/s21platform/auth-proto/auth-proto"
	"github.com/s21platform/gateway-service/internal/repository/grpc"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockAuthClient := NewMockAuth(ctrl)
	MockAuthClient.EXPECT().Login(gomock.Any(), &auth_proto.LoginRequest{
		Username: "testuser",
		Password: "testpass",
	}).Return(&auth_proto.LoginResponse{Jwt: "dummy_token"}, nil)

	grpsC := &grpc.ServiceClients{AuthClient: MockAuthClient}
	r := chi.NewRouter()
	AttachHandlers(r, grpsC)

	//ctx := context.Background()
	t.Run("should get auth", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/auth/login", strings.NewReader(`{"username":"testuser","password":"testpass"}`))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		// Проверка кода ответа
		assert.Equal(t, http.StatusOK, rr.Code, "Expected status code to be 200")

		// Проверка тела ответа
		cookie := rr.Result().Cookies()
		assert.Len(t, cookie, 1, "Expected one cookie")
		assert.Equal(t, "S21SPACE_AUTH_TOKEN", cookie[0].Name, "Cookie name should be 'capy_token'")
		assert.Equal(t, "dummy_token", cookie[0].Value, "Cookie value should be 'dummy_token'")
	})

	t.Run("not found", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/auth/login_fake", strings.NewReader(`{"username":"testuser","password":"testpass"}`))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		// Проверка кода ответа
		assert.Equal(t, http.StatusNotFound, rr.Code, "Expected status code to be 200")
	})

	t.Run("Invalid request data", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/auth/login", strings.NewReader(`{"username":"testuser"}`))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		// Проверка кода ответа
		assert.Equal(t, http.StatusForbidden, rr.Code, "Expected status code to be 200")

	})

	t.Run("Invalid request data 2", func(t *testing.T) {
		MockAuthClient.EXPECT().Login(gomock.Any(), &auth_proto.LoginRequest{
			Username: "testuser",
			Password: "testpass",
		}).Return(nil, errors.New("error"))
		req, err := http.NewRequest("POST", "/auth/login", strings.NewReader(`{"username":"testuser","password":"testpass"}`))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		// Проверка кода ответа
		assert.Equal(t, http.StatusBadRequest, rr.Code, "Expected status code to be 400")
	})
}
