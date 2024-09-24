package auth

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
	"time"
)

type Handler struct {
	aucSrv AuthUsecase
}

func New(aucSrv AuthUsecase) *Handler {
	return &Handler{aucSrv: aucSrv}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Данные введены не полностью", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	jwt, err := h.aucSrv.Login(ctx, data.Username, data.Password)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			http.Error(w, st.Message(), http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "S21SPACE_AUTH_TOKEN",
		Value:    jwt.Jwt,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	})
	w.WriteHeader(http.StatusOK)
	return
}

func AttachAuthRoutes(r chi.Router, handler *Handler) {
	r.Route("/auth", func(authRouter chi.Router) {
		authRouter.Post("/login", handler.Login)
	})
}
