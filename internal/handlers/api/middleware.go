package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	logger_lib "github.com/s21platform/logger-lib"

	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/model"
)

func CheckJWT(next http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
		logger.AddFuncName("CheckJWT")
		cookie, err := r.Cookie("S21SPACE_AUTH_TOKEN")
		if err == nil && cookie.Value != "" {
			logger.Info("older flow (cookie)")
			CheckJWTWithCookie(r, w, next, cfg)
			return
		}
		logger.Info("try new flow (Bearer)")
		accessToken := r.Header.Get("authorization")
		if accessToken == "" {
			logger.Error("access_token is empty")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.Split(accessToken, " ")[1]
		token, err := jwt.ParseWithClaims(tokenStr, &model.ClaimsV2{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(cfg.Platform.AccessSecret), nil
		})
		if err != nil {
			log.Printf("failed to parse token from access_token: %v", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(*model.ClaimsV2)
		if !ok || !token.Valid {
			logger.Error("failed to validation token from access_token")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), config.KeyUsername, claims.Nickname)
		ctx = context.WithValue(ctx, config.KeyUUID, claims.Sub)
		logger.Info(fmt.Sprintf("save to context: %s, %s", claims.Nickname, claims.Sub))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CheckJWTWithCookie(r *http.Request, w http.ResponseWriter, next http.Handler, cfg *config.Config) {
	log.Println("Check jwt for:", r.URL.Path)
	cookie, err := r.Cookie("S21SPACE_AUTH_TOKEN")
	if err != nil {
		log.Println("failed to get cookie value in MW")
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
