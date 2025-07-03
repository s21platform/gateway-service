package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"

	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/model"
)

func CheckJWT(next http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("S21SPACE_AUTH_TOKEN")
		if err == nil {
			CheckJWTWithCookie(r, w, next, cfg)
			return
		}
		accessToken := r.Header.Get("access_token")
		if accessToken == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.Split(accessToken, " ")[1]
		token, err := jwt.ParseWithClaims(tokenStr, &model.ClaimsV2{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(cfg.Platform.Secret), nil
		})
		if err != nil {
			log.Printf("failed to parse token from access_token: %v", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(*model.ClaimsV2)
		if !ok || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), config.KeyUsername, claims.Nickname)
		ctx = context.WithValue(ctx, config.KeyUUID, claims.Sub)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CheckJWTWithCookie(r *http.Request, w http.ResponseWriter, next http.Handler, cfg *config.Config) {
	log.Println("Check jwt for:", r.URL.Path)
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

	token, err := jwt.ParseWithClaims(cookie.Value, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.Platform.Secret), nil
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
	claims, ok := token.Claims.(*model.Claims)
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
	ctx := context.WithValue(r.Context(), config.KeyUsername, claims.Username)
	ctx = context.WithValue(ctx, config.KeyUUID, claims.UUID)
	next.ServeHTTP(w, r.WithContext(ctx))
}
