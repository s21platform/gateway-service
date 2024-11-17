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
	oS OptionService
}

func New(uS UserService, aS AvatarService, oS OptionService) *Handler {
	return &Handler{uS: uS, aS: aS, oS: oS}
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

func (h *Handler) GetAllAvatars(w http.ResponseWriter, r *http.Request) {
	avatars, err := h.aS.GetAvatarsList(r)
	if err != nil {
		log.Printf("get all avatars error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(avatars)
	if err != nil {
		log.Printf("json marshal error: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) DeleteAvatar(w http.ResponseWriter, r *http.Request) {
	deletedAvatar, err := h.aS.RemoveAvatar(r)
	if err != nil {
		log.Printf("delete avatar error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(deletedAvatar)
	if err != nil {
		log.Printf("json marshal error: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetOsById(w http.ResponseWriter, r *http.Request) {
	osInfo, err := h.oS.GetOS(r)
	if err != nil {
		log.Printf("failed to get osInfo: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(osInfo)
	if err != nil {
		log.Printf("json marshal error: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func AttachApiRoutes(r chi.Router, handler *Handler, cfg *config.Config) {
	r.Route("/api", func(apiRouter chi.Router) {
		apiRouter.Use(func(next http.Handler) http.Handler {
			return CheckJWT(next, cfg)
		})

		apiRouter.Get("/profile", handler.MyProfile)
		apiRouter.Post("/avatar", handler.SetAvatar)
		apiRouter.Get("/avatar", handler.GetAllAvatars)
		apiRouter.Delete("/avatar", handler.DeleteAvatar)
		apiRouter.Get("/option/os", handler.GetOsById)
	})
}
