package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	auth_proto "github.com/s21platform/auth-proto/auth-proto"
	"github.com/s21platform/gateway-service/internal/config"
	grpc2 "github.com/s21platform/gateway-service/internal/repository/grpc"
	"github.com/s21platform/gateway-service/internal/useCase/RESTHandlers"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	cfg := config.MustLoad()

	authConn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.Auth.Host, cfg.Auth.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to auth service: %v", err)
	}
	defer authConn.Close()
	authClient := auth_proto.NewAuthServiceClient(authConn)

	clients := &grpc2.GrpcClients{
		AuthClient: authClient,
	}
	r := chi.NewRouter()
	RESTHandlers.AttachHandlers(r, clients)
	fmt.Println(fmt.Sprintf(":%s", cfg.Service.Port))
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.Service.Port), r)
}
