package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/status"

	logger_lib "github.com/s21platform/logger-lib"

	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/model"
)

type Handler struct {
	secretToken string
	aucSrv      Usecase
}

func New(cfg *config.Config, aucSrv Usecase) *Handler {
	return &Handler{aucSrv: aucSrv, secretToken: cfg.Platform.Secret}
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

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "S21SPACE_AUTH_TOKEN",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
}

func (h *Handler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("CheckAuth")

	cookie, err := r.Cookie("S21SPACE_AUTH_TOKEN")
	if err != nil {
		logger.Error("failed to get cookie value")
		log.Println("failed to get cookie value")
		http.SetCookie(w, &http.Cookie{
			Name:     "S21SPACE_AUTH_TOKEN",
			Value:    "",
			MaxAge:   -1,
			HttpOnly: true,
		})
		w.WriteHeader(http.StatusUnauthorized)
		msg, err := prepareResponse("failed to get cookie value", false)
		if err != nil {
			logger.Error(fmt.Sprintf("failed make response: %s", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("failed make response"))
			return
		}
		_, _ = w.Write(msg)
		return
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(h.secretToken), nil
	})
	if err != nil {
		logger.Error(fmt.Sprintf("failed to parse token: %v", err))
		log.Printf("failed to parse token: %v", err)
		http.SetCookie(w, &http.Cookie{
			Name:     "S21SPACE_AUTH_TOKEN",
			Value:    "",
			MaxAge:   -1,
			HttpOnly: true,
		})
		w.WriteHeader(http.StatusUnauthorized)
		msg, err := prepareResponse("failed to parse token", false)
		if err != nil {
			logger.Error(fmt.Sprintf("failed make response: %s", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("failed make response"))
			return
		}
		_, _ = w.Write(msg)
		return
	}
	if !token.Valid {
		http.SetCookie(w, &http.Cookie{
			Name:     "S21SPACE_AUTH_TOKEN",
			Value:    "",
			MaxAge:   -1,
			HttpOnly: true,
		})
		w.WriteHeader(http.StatusUnauthorized)
		msg, err := prepareResponse("token not valid", true)
		if err != nil {
			logger.Error(fmt.Sprintf("failed make response: %s", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("failed make response"))
			return
		}
		_, _ = w.Write(msg)
		return
	}
	w.WriteHeader(http.StatusOK)
	msg, err := prepareResponse("", true)
	if err != nil {
		logger.Error(fmt.Sprintf("failed make response: %s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed make response"))
		return
	}
	_, _ = w.Write(msg)
}

func prepareResponse(message string, isAuth bool) ([]byte, error) {
	resp := model.CheckAuth{
		IsAuth: isAuth,
		Error:  message,
	}
	msg, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (h *Handler) CheckEmailAvailability(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("CheckEmailAvailability")

	result, err := h.aucSrv.CheckEmailAvailability(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to check email: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(&result)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to json marshal: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) SendUserVerificationCode(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("SendUserVerificationCode")

	result, err := h.aucSrv.SendUserVerificationCode(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to send user verification code: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(&result)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to json marshal: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("RegisterUser")

	resp, err := h.aucSrv.RegisterUser(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to register user: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(&resp)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to json marshal: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) LoginV2(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("LoginV2")

	resp, err := h.aucSrv.LoginV2(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to login: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(&model.LoginV2Response{AccessToken: resp.AccessToken})
	if err != nil {
		logger.Error(fmt.Sprintf("failed to json marshal: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    resp.RefreshToken,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)

	logger.Info("OK")
}

func AttachAuthRoutes(r chi.Router, handler *Handler) {
	r.Route("/auth", func(authRouter chi.Router) {
		authRouter.Post("/login", handler.Login)
		authRouter.Get("/check-auth", handler.CheckAuth)
		authRouter.Get("/logout", handler.Logout)
		authRouter.Get("/check-email", handler.CheckEmailAvailability)
		authRouter.Post("/send-code", handler.SendUserVerificationCode)
		authRouter.Post("/register-user", handler.RegisterUser)
		authRouter.Post("/v2/login", handler.LoginV2)
	})
}
