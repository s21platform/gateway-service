package auth

import (
	"context"
	"encoding/json"
	"fmt"
	auth_proto "github.com/s21platform/auth-proto/auth-proto"
	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/repository/auth"
	"google.golang.org/grpc/status"
	"net/http"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetAuth(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	//authService := auth.NewAuthServiceClient(cfg)
	return func(w http.ResponseWriter, r *http.Request) {
		data := &LoginData{}
		err := json.NewDecoder(r.Body).Decode(data)
		if err != nil {
			fmt.Println("some error for decode data")
			return
		}
		authService := auth.NewAuthServiceClient(cfg)
		client := auth_proto.NewAuthServiceClient(authService.Conn)
		request := auth_proto.LoginRequest{Username: data.Username, Password: data.Password}
		response, err := client.Login(context.Background(), &request)
		if err != nil {
			if statusError, ok := status.FromError(err); ok {
				http.Error(w, statusError.Message(), int(statusError.Code()))
				return
			}
		}
		http.SetCookie(w, &http.Cookie{
			Name:  "capy_token",
			Value: response.Jwt,
		})
		return
	}
}
