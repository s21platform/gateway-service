//go:build !test

package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/handlers/api"
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
	apiHandlers := api.New()

	r := chi.NewRouter()
	authhandler.AttachAuthRoutes(r, authHandlers)
	api.AttachApiRoutes(r, apiHandlers, cfg)

	fmt.Println(fmt.Sprintf(":%s", cfg.Service.Port))

	http.ListenAndServe(fmt.Sprintf(":%s", cfg.Service.Port), r)
}
