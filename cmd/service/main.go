//go:build !test

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/s21platform/metrics-lib/pkg"

	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/handlers/api"
	authhandler "github.com/s21platform/gateway-service/internal/handlers/auth"
	"github.com/s21platform/gateway-service/internal/middlewares"
	"github.com/s21platform/gateway-service/internal/rpc/auth"
	"github.com/s21platform/gateway-service/internal/rpc/avatar"
	"github.com/s21platform/gateway-service/internal/rpc/friends"
	"github.com/s21platform/gateway-service/internal/rpc/notification"
	"github.com/s21platform/gateway-service/internal/rpc/option"
	"github.com/s21platform/gateway-service/internal/rpc/society"
	"github.com/s21platform/gateway-service/internal/rpc/user"
	authusecase "github.com/s21platform/gateway-service/internal/useCase/auth"
	avatarusecase "github.com/s21platform/gateway-service/internal/useCase/avatar"
	friendsusecase "github.com/s21platform/gateway-service/internal/useCase/friends"
	notificationusecase "github.com/s21platform/gateway-service/internal/useCase/notification"
	optionusecase "github.com/s21platform/gateway-service/internal/useCase/option"
	societyusecase "github.com/s21platform/gateway-service/internal/useCase/society"
	userusecase "github.com/s21platform/gateway-service/internal/useCase/user"
)

func main() {
	cfg := config.MustLoad()

	metrics, err := pkg.NewMetrics(cfg.Metrics.Host, cfg.Metrics.Port, "gateway", cfg.Platform.Env)
	if err != nil {
		log.Fatalf("failed to init metrics: %v", err)
	}
	defer metrics.Disconnect()

	// rpc clients
	authClient := auth.NewService(cfg)
	userClient := user.NewService(cfg)
	avatarClient := avatar.New(cfg)
	notificationClient := notification.New(cfg)
	friendsClient := friends.NewService(cfg)
	optionClient := option.New(cfg)
	societyClient := society.NewService(cfg)

	// usecases declaration
	authUseCase := authusecase.New(authClient)
	userUsecase := userusecase.New(userClient)
	avatarUsecase := avatarusecase.New(avatarClient)
	notificationUsecase := notificationusecase.New(notificationClient)
	friendsUseCase := friendsusecase.New(friendsClient)
	optionUsecase := optionusecase.New(optionClient)
	societyUseCase := societyusecase.New(societyClient)

	// handlers declaration
	authHandlers := authhandler.New(authUseCase)
	apiHandlers := api.New(userUsecase, avatarUsecase, notificationUsecase, friendsUseCase, optionUsecase, societyUseCase)

	r := chi.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return middlewares.MetricMiddleware(next, metrics)
	})

	authhandler.AttachAuthRoutes(r, authHandlers)
	api.AttachApiRoutes(r, apiHandlers, cfg)

	log.Println("Server starting...")

	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Service.Port), r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
