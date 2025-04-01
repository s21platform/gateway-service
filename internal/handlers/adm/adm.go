package adm

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	api "github.com/s21platform/staff-service/pkg/staff/v0"
)

type Handler struct {
	sC StaffClient
}

func New(sC StaffClient) *Handler {
	return &Handler{sC: sC}
}

func (h *Handler) StaffLogin(w http.ResponseWriter, r *http.Request) {
	var loginRequest api.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	response, err := h.sC.StaffLogin(r.Context(), &loginRequest)
	if err != nil {
		log.Println("failed to login", err)
		http.Error(w, "failed to login", http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func AttachAdmRoutes(r chi.Router, handler *Handler) {
	r.Route("/adm", func(apiRouter chi.Router) {
		apiRouter.Use(func(next http.Handler) http.Handler {
			return CheckJWT(next)
		})

		apiRouter.Post("/auth/login", handler.StaffLogin)
	})
}
