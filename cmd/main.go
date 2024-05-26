//go:build !test

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
	// load environments to Config struct
	cfg := config.MustLoad()

	// Create connection to Auth service by gRPC
	authConn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.Auth.Host, cfg.Auth.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to auth service: %v", err)
	}
	defer authConn.Close()

	// Make struct for storage all Clients connection of this service
	clients := &grpc2.ServiceClients{
		AuthClient: auth_proto.NewAuthServiceClient(authConn),
	}

	// Create New Router for manage url endpoints
	r := chi.NewRouter()

	// Register all REST Handlers
	RESTHandlers.AttachHandlers(r, clients)

	// Print Port from environments for check load env
	fmt.Println(fmt.Sprintf(":%s", cfg.Service.Port))

	// Listen port waiting for requests
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.Service.Port), r)
}
