package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/s21platform/gateway-service/internal/config"
	logger_lib "github.com/s21platform/logger-lib"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc/status"
)

type Handler struct {
	aucSrv Usecase
}

func New(aucSrv Usecase) *Handler {
	return &Handler{aucSrv: aucSrv}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("Login")
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &data); err != nil {
		logger.Error(err.Error())
		http.Error(w, "Данные введены не полностью", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	jwt, err := h.aucSrv.Login(ctx, data.Username, data.Password)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			logger.Error(st.Message())
			http.Error(w, st.Message(), http.StatusForbidden)
			return
		}
		logger.Error(err.Error())
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
	logger.Info("OK")
}

func AttachAuthRoutes(r chi.Router, handler *Handler) {
	r.Route("/auth", func(authRouter chi.Router) {
		authRouter.Post("/login", handler.Login)
	})
}
