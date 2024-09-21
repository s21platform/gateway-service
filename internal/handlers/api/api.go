package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	return
}

func CheckJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Извлекаем JWT токен из заголовка Authorization
		_, err := r.Cookie("S21SPACE_AUTH_TOKEN")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		//token := strings.TrimPrefix(authHeader, "Bearer ")

		// Здесь вы должны добавить логику валидации JWT
		//valid, err := ValidateJWT(token)
		//if err != nil || !valid {
		//	http.Error(w, "Invalid token", http.StatusUnauthorized)
		//	return
		//}

		// Если JWT валиден, передаем запрос дальше
		next.ServeHTTP(w, r)
	})
}

func AttachApiRoutes(r chi.Router, handler *Handler) {
	r.Group(func(r chi.Router) {
		r.Use(CheckJWT)
		r.Route("/api", func(r chi.Router) {
			r.Get("/test", handler.Test)
		})
	})
}
