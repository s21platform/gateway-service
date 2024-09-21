//go:build !test

package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/s21platform/gateway-service/internal/config"
	authhandler "github.com/s21platform/gateway-service/internal/handlers/auth"
	"github.com/s21platform/gateway-service/internal/rpc/auth"
	authusecase "github.com/s21platform/gateway-service/internal/useCase/auth"
	"net/http"
)

func main() {
	cfg := config.MustLoad()

	// rpc clients
	authClient := auth.NewService(cfg)

	// usecases declaration
	authUseCase := authusecase.New(authClient)

	// handlers declaration
	authHandlers := authhandler.New(authUseCase)

	r := chi.NewRouter()
	authhandler.AttachAuthRoutes(r, authHandlers)

	fmt.Println(fmt.Sprintf(":%s", cfg.Service.Port))
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.Service.Port), r)
}
