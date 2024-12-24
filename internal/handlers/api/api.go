package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	logger_lib "github.com/s21platform/logger-lib"

	"github.com/go-chi/chi/v5"

	"github.com/s21platform/gateway-service/internal/config"
)

type Handler struct {
	uS  UserService
	aS  AvatarService
	nS  NotificationService
	fs  FriendsService
	oS  OptionService
	sS  SocietyService
	srS SearchService
}

func New(uS UserService, aS AvatarService, nS NotificationService, fS FriendsService, oS OptionService, sS SocietyService, srS SearchService) *Handler {
	return &Handler{uS: uS, aS: aS, nS: nS, fs: fS, oS: oS, sS: sS, srS: srS}
}

func (h *Handler) MyProfile(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("MyProfile")

	resp, err := h.uS.GetInfoByUUID(r)
	if err != nil {
		logger.Error(fmt.Sprintf("get info by uuid error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(resp)
	if err != nil {
		logger.Error(fmt.Sprintf("json marshal error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(jsn)
}

func (h *Handler) PeerInfo(w http.ResponseWriter, r *http.Request) {
	resp, err := h.uS.GetPeerInfo(r)
	if err != nil {
		log.Printf("get peer info error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsn, err := json.Marshal(resp)
	if err != nil {
		log.Printf("json marshal error: %v", err)
	}
	_, _ = w.Write(jsn)
}

func (h *Handler) SetAvatar(w http.ResponseWriter, r *http.Request) {
	resp, err := h.aS.UploadAvatar(r)
	if err != nil {
		log.Printf("upload avatar error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(resp)
	if err != nil {
		log.Printf("json marshal error: %v", err)
	}
	_, _ = w.Write(jsn)
}

func (h *Handler) GetAllAvatars(w http.ResponseWriter, r *http.Request) {
	avatars, err := h.aS.GetAvatarsList(r)
	if err != nil {
		log.Printf("get all avatars error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(avatars)
	if err != nil {
		log.Printf("json marshal error: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) DeleteAvatar(w http.ResponseWriter, r *http.Request) {
	deletedAvatar, err := h.aS.RemoveAvatar(r)
	if err != nil {
		log.Printf("delete avatar error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(deletedAvatar)
	if err != nil {
		log.Printf("json marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) CountNotifications(w http.ResponseWriter, r *http.Request) {
	result, err := h.nS.GetCountNotification(r)
	if err != nil {
		log.Printf("get count notification error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(result)
	if err != nil {
		log.Printf("json marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetNotifications(w http.ResponseWriter, r *http.Request) {
	result, err := h.nS.GetNotification(r)
	if err != nil {
		log.Printf("get notification error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsn, err := json.Marshal(result)
	if err != nil {
		log.Printf("json marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetCountFriends(w http.ResponseWriter, r *http.Request) {
	result, err := h.fs.GetCountFriends(r)
	if err != nil {
		log.Printf("failed to get friends error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsn, err := json.Marshal(result)
	if err != nil {
		log.Printf("failed to json marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) SetFriends(w http.ResponseWriter, r *http.Request) {
	result, err := h.fs.SetFriends(r)
	if err != nil {
		log.Printf("failed to set friends error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsn, err := json.Marshal(result)
	if err != nil {
		log.Printf("failed to json marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsn)
}

func (h *Handler) RemoveFriends(w http.ResponseWriter, r *http.Request) {
	result, err := h.fs.RemoveFriends(r)
	if err != nil {
		log.Printf("failed to remove friends error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsn, err := json.Marshal(result)
	if err != nil {
		log.Printf("failed to json marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("json: ", string(jsn))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetUsersWithLimit(w http.ResponseWriter, r *http.Request) {
	result, err := h.srS.GetUserWithLimit(r)
	if err != nil {
		log.Printf("failed to get users with limit error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsn, err := json.Marshal(result)
	if err != nil {
		log.Printf("failed to json marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetOsBySearchName(w http.ResponseWriter, r *http.Request) {
	osList, err := h.oS.GetOsList(r)
	if err != nil {
		log.Printf("failed to get os list: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(osList)
	if err != nil {
		log.Printf("json marshal error: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetWorkPlaceBySearchName(w http.ResponseWriter, r *http.Request) {
	osList, err := h.oS.GetWorkPlaceList(r)
	if err != nil {
		log.Printf("failed to get workplace list: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(osList)
	if err != nil {
		log.Printf("json marshal error: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetStudyPlaceBySearchName(w http.ResponseWriter, r *http.Request) {
	osList, err := h.oS.GetStudyPlaceList(r)
	if err != nil {
		log.Printf("failed to get study place list: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(osList)
	if err != nil {
		log.Printf("json marshal error: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetHobbyBySearchName(w http.ResponseWriter, r *http.Request) {
	osList, err := h.oS.GetHobbyList(r)
	if err != nil {
		log.Printf("failed to get hobby list: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(osList)
	if err != nil {
		log.Printf("json marshal error: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetSkillBySearchName(w http.ResponseWriter, r *http.Request) {
	osList, err := h.oS.GetSkillList(r)
	if err != nil {
		log.Printf("failed to get skill list: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(osList)
	if err != nil {
		log.Printf("json marshal error: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetCityBySearchName(w http.ResponseWriter, r *http.Request) {
	osList, err := h.oS.GetCityList(r)
	if err != nil {
		log.Printf("failed to get city list: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(osList)
	if err != nil {
		log.Printf("json marshal error: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetSocietyDirectionBySearchName(w http.ResponseWriter, r *http.Request) {
	osList, err := h.oS.GetSocietyDirectionList(r)
	if err != nil {
		log.Printf("failed to get society direction list: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(osList)
	if err != nil {
		log.Printf("json marshal error: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("UpdateProfile")

	resp, err := h.uS.UpdateProfileInfo(r)
	if err != nil {
		logger.Error(fmt.Sprintf("update profile info error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsn, err := json.Marshal(resp)
	if err != nil {
		logger.Error(fmt.Sprintf("json marshal error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) CreateSociety(w http.ResponseWriter, r *http.Request) {
	result, err := h.sS.CreateSociety(r)
	if err != nil {
		log.Printf("create society error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsn, err := json.Marshal(result)
	if err != nil {
		log.Printf("json marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("json: ", string(jsn))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetAccessLevel(w http.ResponseWriter, r *http.Request) {
	result, err := h.sS.GetAccessLevel(r)
	if err != nil {
		log.Printf("get access level error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsn, err := json.Marshal(result)
	if err != nil {
		log.Printf("json marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetSocietyInfo(w http.ResponseWriter, r *http.Request) {
	result, err := h.sS.GetSocietyInfo(r)
	if err != nil {
		log.Printf("get society info error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsn, err := json.Marshal(result)
	if err != nil {
		log.Printf("json marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func AttachApiRoutes(r chi.Router, handler *Handler, cfg *config.Config) {
	r.Route("/api", func(apiRouter chi.Router) {
		apiRouter.Use(func(next http.Handler) http.Handler {
			return CheckJWT(next, cfg)
		})

		apiRouter.Get("/profile", handler.MyProfile)
		apiRouter.Put("/profile", handler.UpdateProfile)
		apiRouter.Post("/avatar", handler.SetAvatar)
		apiRouter.Get("/avatar", handler.GetAllAvatars)
		apiRouter.Delete("/avatar", handler.DeleteAvatar)
		apiRouter.Get("/notification/count", handler.CountNotifications)
		apiRouter.Get("/notification", handler.GetNotifications)
		apiRouter.Get("/friends/counts", handler.GetCountFriends)
		apiRouter.Get("/option/os", handler.GetOsBySearchName)
		apiRouter.Get("/option/workplace", handler.GetWorkPlaceBySearchName)
		apiRouter.Get("/option/study-place", handler.GetStudyPlaceBySearchName)
		apiRouter.Get("/option/hobby", handler.GetHobbyBySearchName)
		apiRouter.Get("/option/skill", handler.GetSkillBySearchName)
		apiRouter.Get("/option/city", handler.GetCityBySearchName)
		apiRouter.Get("/option/society-direction", handler.GetSocietyDirectionBySearchName)
		apiRouter.Post("/society", handler.CreateSociety)
		apiRouter.Get("/society/access", handler.GetAccessLevel)
		apiRouter.Get("/society", handler.GetSocietyInfo)
		apiRouter.Post("/user", handler.SetFriends)
		apiRouter.Delete("/user", handler.RemoveFriends)
		apiRouter.Get("/peer/{uuid}", handler.PeerInfo)
		apiRouter.Get("/search", handler.GetUsersWithLimit)

	})
}
