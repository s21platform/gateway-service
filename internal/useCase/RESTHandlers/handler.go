package RESTHandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	auth_proto "github.com/s21platform/auth-proto/auth-proto"
	grpc2 "github.com/s21platform/gateway-service/internal/repository/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
)

func AttachHandlers(r chi.Router, clients *grpc2.ServiceClients) {
	// Register all endpoints for /auth/
	r.Route("/auth", func(authRouter chi.Router) {
		authRouter.Post("/login", GetAuth(clients))
	})
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

func GetAuth(in Auth) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data := &LoginData{}
		err := json.NewDecoder(r.Body).Decode(data)
		if err != nil || data.Username == "" || data.Password == "" {
			fmt.Println("some error for decode data")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(LoginResponse{
				Message: "Необходимо заполнить все поля",
			})
			return
		}
		ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("trace-id", "-test-"))
		request := auth_proto.LoginRequest{Username: data.Username, Password: data.Password}
		response, err := in.Login(ctx, &request)
		if err != nil {
			if statusError, ok := status.FromError(err); ok {
				switch statusError.Code() {
				case codes.InvalidArgument:
					w.WriteHeader(http.StatusForbidden)
					json.NewEncoder(w).Encode(LoginResponse{
						Message: statusError.Message(),
					})
				default:
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(LoginResponse{
						Message: "Unknown status code",
					})
				}
				log.Println(int(statusError.Code()), statusError.Message())
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(LoginResponse{
				Message: "Unknown error",
			})
			log.Printf("Error after school service: %v", err)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:  "S21SPACE_AUTH_TOKEN",
			Value: response.Jwt,
		})
		w.WriteHeader(http.StatusOK)
		return
	}
}
