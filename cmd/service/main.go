//go:build !test

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	logger_lib "github.com/s21platform/logger-lib"
	"github.com/s21platform/metrics-lib/pkg"

	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/handlers/adm"
	"github.com/s21platform/gateway-service/internal/handlers/api"
	authhandler "github.com/s21platform/gateway-service/internal/handlers/auth"
	"github.com/s21platform/gateway-service/internal/middlewares"
	"github.com/s21platform/gateway-service/internal/rpc/advert"
	"github.com/s21platform/gateway-service/internal/rpc/auth"
	"github.com/s21platform/gateway-service/internal/rpc/avatar"
	"github.com/s21platform/gateway-service/internal/rpc/chat"
	"github.com/s21platform/gateway-service/internal/rpc/community"
	"github.com/s21platform/gateway-service/internal/rpc/feed"

	//"github.com/s21platform/gateway-service/internal/rpc/friends"
	"github.com/s21platform/gateway-service/internal/rpc/materials"
	"github.com/s21platform/gateway-service/internal/rpc/notification"
	"github.com/s21platform/gateway-service/internal/rpc/option"
	"github.com/s21platform/gateway-service/internal/rpc/search"
	"github.com/s21platform/gateway-service/internal/rpc/society"
	"github.com/s21platform/gateway-service/internal/rpc/staff"
	"github.com/s21platform/gateway-service/internal/rpc/user"
	advertusecase "github.com/s21platform/gateway-service/internal/useCase/advert"
	authusecase "github.com/s21platform/gateway-service/internal/useCase/auth"
	avatarusecase "github.com/s21platform/gateway-service/internal/useCase/avatar"
	chatusecase "github.com/s21platform/gateway-service/internal/useCase/chat"
	communityusecase "github.com/s21platform/gateway-service/internal/useCase/community"
	feedusecase "github.com/s21platform/gateway-service/internal/useCase/feed"
	materialsusecase "github.com/s21platform/gateway-service/internal/useCase/materials"

	//friendsusecase "github.com/s21platform/gateway-service/internal/useCase/friends"
	notificationusecase "github.com/s21platform/gateway-service/internal/useCase/notification"
	optionusecase "github.com/s21platform/gateway-service/internal/useCase/option"
	searchusecase "github.com/s21platform/gateway-service/internal/useCase/search"
	societyusecase "github.com/s21platform/gateway-service/internal/useCase/society"
	staffusecase "github.com/s21platform/gateway-service/internal/useCase/staff"
	userusecase "github.com/s21platform/gateway-service/internal/useCase/user"
)

func main() {
	cfg := config.MustLoad()

	metrics, err := pkg.NewMetrics(cfg.Metrics.Host, cfg.Metrics.Port, "gateway", cfg.Platform.Env)
	if err != nil {
		log.Fatalf("failed to init metrics: %v", err)
	}
	defer metrics.Disconnect()

	logger := logger_lib.New(cfg.Logger.Host, cfg.Logger.Port, cfg.Service.Name, cfg.Platform.Env)

	// rpc clients
	authClient := auth.NewService(cfg)
	userClient := user.NewService(cfg)
	avatarClient := avatar.New(cfg)
	notificationClient := notification.New(cfg)
	//friendsClient := friends.NewService(cfg)
	optionClient := option.New(cfg)
	societyClient := society.NewService(cfg)
	searchClient := search.NewService(cfg)
	chatClient := chat.NewService(cfg)
	advertClient := advert.New(cfg)
	feedClient := feed.New(cfg)
	staffClient := staff.New(cfg)
	materialsClient := materials.NewService(cfg)
	CommunityClient := community.New(cfg)

	// usecases declaration
	authUseCase := authusecase.New(authClient)
	userUsecase := userusecase.New(userClient)
	avatarUsecase := avatarusecase.New(avatarClient)
	notificationUsecase := notificationusecase.New(notificationClient)
	//friendsUseCase := friendsusecase.New(friendsClient)
	optionUsecase := optionusecase.New(optionClient)
	societyUseCase := societyusecase.New(societyClient)
	searchUseCase := searchusecase.New(searchClient)
	chatUseCase := chatusecase.New(chatClient)
	advertUseCase := advertusecase.New(advertClient)
	feedUseCase := feedusecase.New(feedClient)
	staffUseCase := staffusecase.New(staffClient)
	materialsUseCase := materialsusecase.New(materialsClient)
	communityUseCase := communityusecase.New(CommunityClient)

	// handlers declaration
	authHandlers := authhandler.New(cfg, authUseCase)
	apiHandlers := api.New(userUsecase, avatarUsecase, notificationUsecase, optionUsecase, societyUseCase, searchUseCase, chatUseCase, advertUseCase, feedUseCase, materialsUseCase, communityUseCase)
	admHandlers := adm.New(staffUseCase)

	r := chi.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return middlewares.MetricMiddleware(next, metrics)
	})
	r.Use(func(next http.Handler) http.Handler {
		return middlewares.LoggerMiddleware(next, logger)
	})

	authhandler.AttachAuthRoutes(r, authHandlers)
	api.AttachApiRoutes(r, apiHandlers, cfg)
	adm.AttachAdmRoutes(r, admHandlers)

	log.Println("Server starting...")

	if err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.Service.Port), r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
