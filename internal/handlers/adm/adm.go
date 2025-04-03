package adm

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

func (h *Handler) ListStaff(w http.ResponseWriter, r *http.Request) {
	req := &api.ListStaffRequest{
		Page:     1,  // значение по умолчанию
		PageSize: 10, // значение по умолчанию
	}

	// Обработка page
	if page := r.URL.Query().Get("page"); page != "" {
		pageInt, err := strconv.ParseInt(page, 10, 32)
		if err != nil {
			http.Error(w, "invalid page parameter", http.StatusBadRequest)
			return
		}
		req.Page = int32(pageInt)
	}

	// Обработка page_size
	if pageSize := r.URL.Query().Get("page_size"); pageSize != "" {
		pageSizeInt, err := strconv.ParseInt(pageSize, 10, 32)
		if err != nil {
			http.Error(w, "invalid page_size parameter", http.StatusBadRequest)
			return
		}
		req.PageSize = int32(pageSizeInt)
	}

	// Обработка search_term
	if searchTerm := r.URL.Query().Get("search_term"); searchTerm != "" {
		req.SearchTerm = &searchTerm
	}

	// Обработка role_id
	if roleID := r.URL.Query().Get("role_id"); roleID != "" {
		roleIDInt, err := strconv.ParseInt(roleID, 10, 32)
		if err != nil {
			http.Error(w, "invalid role_id parameter", http.StatusBadRequest)
			return
		}
		roleIDInt32 := int32(roleIDInt)
		req.RoleId = &roleIDInt32
	}

	staffList, err := h.sC.ListStaff(r.Context(), req)
	if err != nil {
		http.Error(w, "failed to list staff", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(staffList); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func AttachAdmRoutes(r chi.Router, handler *Handler) {
	r.Route("/adm", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", handler.StaffLogin)
		})
		r.With(CheckJWT).Route("/staff", func(r chi.Router) {
			r.Post("/", handler.CreateStaff)
			r.Get("/list", handler.ListStaff)
		})
	})
}
