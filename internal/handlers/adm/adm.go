package adm

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	logger_lib "github.com/s21platform/logger-lib"

	"github.com/s21platform/gateway-service/internal/config"
)

type Handler struct {
	sS StaffService
}

func New(sS StaffService) *Handler {
	return &Handler{sS: sS}
}

func (h *Handler) StaffLogin(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("StaffLogin")

	result, err := h.sS.StaffLogin(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to login: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(result)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to marshal response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) CreateStaff(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("CreateStaff")

	result, err := h.sS.CreateStaff(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create staff: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(result)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to marshal response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) ListStaff(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("ListStaff")

	result, err := h.sS.ListStaff(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to list staff: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(result)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to marshal response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetStaff(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("GetStaff")

	staffID := chi.URLParam(r, "uuid")
	result, err := h.sS.GetStaff(r, staffID)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get staff: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(result)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to marshal response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func AttachAdmRoutes(r chi.Router, handler *Handler) {
	r.Route("/adm", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", handler.StaffLogin)
		})
		r.With(CheckJWT).Route("/staff", func(r chi.Router) {
			r.Post("/", handler.CreateStaff)
			r.Get("/list", handler.ListStaff)
			r.Get("/{uuid}", handler.GetStaff)
		})
	})
}
