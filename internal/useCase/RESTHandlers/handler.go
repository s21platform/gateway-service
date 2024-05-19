package RESTHandlers

import (
	"github.com/go-chi/chi"
	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/useCase/auth"
)

func AttachHandlers(r chi.Router, cfg *config.Config) {
	r.Route("/auth", func(authRouter chi.Router) {
		authRouter.Post("/login", auth.GetAuth(cfg))
	})
}
