package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/s21platform/gateway-service/internal/config"
)

type Handler struct {
	uS UserService
	aS AvatarService
}

func New(uS UserService, aS AvatarService) *Handler {
	return &Handler{uS: uS, aS: aS}
}

func (h *Handler) MyProfile(w http.ResponseWriter, r *http.Request) {
	resp, err := h.uS.GetInfoByUUID(r.Context())
	if err != nil {
		log.Printf("get info by uuid error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(resp)
	if err != nil {
		log.Printf("json marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(jsn)
}

func (h *Handler) SetAvatar(w http.ResponseWriter, r *http.Request) {
	resp, err := h.aS.UploadAvatar(r)
	if err != nil {
		log.Printf("upload avatar error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(resp)
	if err != nil {
		log.Printf("json marshal error: %v", err)
	}
	_, _ = w.Write(jsn)
}

func AttachApiRoutes(r chi.Router, handler *Handler, cfg *config.Config) {
	r.Route("/api", func(apiRouter chi.Router) {
		apiRouter.Use(func(next http.Handler) http.Handler {
			return CheckJWT(next, cfg)
		})

		apiRouter.Get("/profile", handler.MyProfile)
		apiRouter.Post("/avatar", handler.SetAvatar)
	})
}
