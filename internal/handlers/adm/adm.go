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

func (h *Handler) CreateStaff(w http.ResponseWriter, r *http.Request) {
	var req api.CreateStaffRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	staff, err := h.sC.CreateStaff(r.Context(), &req)
	if err != nil {
		http.Error(w, "failed to create staff", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(staff); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func AttachAdmRoutes(r chi.Router, handler *Handler) {
	r.Route("/adm", func(apiRouter chi.Router) {
		apiRouter.Use(func(next http.Handler) http.Handler {
			return CheckJWT(next)
		})

		apiRouter.Post("/auth/login", handler.StaffLogin)
		apiRouter.Post("/staff", handler.CreateStaff)
	})
}
