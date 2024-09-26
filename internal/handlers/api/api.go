package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"github.com/s21platform/gateway-service/internal/config"
	"log"
	"net/http"
)

type Handler struct {
	uS UserService
}

func New(uS UserService) *Handler {
	return &Handler{uS: uS}
}

func (h *Handler) Test(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username")
	w.Write([]byte(fmt.Sprintf("Hello, %s!", username)))
	return
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
	w.Write(jsn)
	return
}

type Claims struct {
	UUID        string `json:"uid"`
	Username    string `json:"username"`
	Role        string `json:"role"`
	AccessToken string `json:"accessToken"`
	Exp         int64  `json:"exp"`
	jwt.RegisteredClaims
}

func CheckJWT(next http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("S21SPACE_AUTH_TOKEN")
		if err != nil {
			log.Println("failed to get cookie value")
			http.SetCookie(w, &http.Cookie{
				Name:     "S21SPACE_AUTH_TOKEN",
				Value:    "",
				MaxAge:   -1,
				HttpOnly: true,
			})
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			log.Println("token:", cfg.Service.Secret)
			return []byte(cfg.Service.Secret), nil
		})
		if err != nil {
			log.Printf("failed to parse token: %v", err)
			http.SetCookie(w, &http.Cookie{
				Name:     "S21SPACE_AUTH_TOKEN",
				Value:    "",
				MaxAge:   -1,
				HttpOnly: true,
			})
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			http.SetCookie(w, &http.Cookie{
				Name:     "S21SPACE_AUTH_TOKEN",
				Value:    "",
				MaxAge:   -1,
				HttpOnly: true,
			})
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "username", claims.Username)
		ctx = context.WithValue(ctx, "uuid", claims.UUID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AttachApiRoutes(r chi.Router, handler *Handler, cfg *config.Config) {
	r.Group(func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return CheckJWT(next, cfg)
		})

		r.Route("/api", func(r chi.Router) {
			r.Get("/profile", handler.MyProfile)
		})
	})
}
