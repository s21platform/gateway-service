package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	logger_lib "github.com/s21platform/logger-lib"

	"github.com/s21platform/gateway-service/internal/config"
)

type Handler struct {
	uS  UserService
	aS  AvatarService
	nS  NotificationService
	fS  FriendsService
	oS  OptionService
	sS  SocietyService
	srS SearchService
	cS  ChatService
	adS AdvertService
}

func New(uS UserService, aS AvatarService, nS NotificationService, fS FriendsService, oS OptionService, sS SocietyService, srS SearchService, cS ChatService, adS AdvertService) *Handler {
	return &Handler{uS: uS, aS: aS, nS: nS, fS: fS, oS: oS, sS: sS, srS: srS, cS: cS, adS: adS}
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

func (h *Handler) SetUserAvatar(w http.ResponseWriter, r *http.Request) {
	resp, err := h.aS.UploadUserAvatar(r)
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

func (h *Handler) GetAllUserAvatars(w http.ResponseWriter, r *http.Request) {
	avatars, err := h.aS.GetUserAvatarsList(r)
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

func (h *Handler) DeleteUserAvatar(w http.ResponseWriter, r *http.Request) {
	deletedAvatar, err := h.aS.RemoveUserAvatar(r)
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

func (h *Handler) SetSocietyAvatar(w http.ResponseWriter, r *http.Request) {
	resp, err := h.aS.UploadSocietyAvatar(r)
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

func (h *Handler) GetAllSocietyAvatars(w http.ResponseWriter, r *http.Request) {
	avatars, err := h.aS.GetSocietyAvatarsList(r)
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

func (h *Handler) DeleteSocietyAvatar(w http.ResponseWriter, r *http.Request) {
	deletedAvatar, err := h.aS.RemoveSocietyAvatar(r)
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

func (h *Handler) MarkNotificationAsRead(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("MarkNotificationAsRead")

	if _, err := h.nS.MarkNotificationsAsRead(r); err != nil {
		logger.Error(fmt.Sprintf("failed to mark notification as read: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("{}\n"))
}

func (h *Handler) GetCountFriends(w http.ResponseWriter, r *http.Request) {
	result, err := h.fS.GetCountFriends(r)
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
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("SetFriends")
	result, err := h.fS.SetFriends(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to set friends error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsn, err := json.Marshal(result)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to json marshal error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsn)
}

func (h *Handler) RemoveFriends(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("RemoveFriends")
	result, err := h.fS.RemoveFriends(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to remove friends error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsn, err := json.Marshal(result)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to json marshal error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("json: ", string(jsn))
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
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("CreateSociety")
	result, err := h.sS.CreateSociety(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create society error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsn, err := json.Marshal(result)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to json marshal error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetSocietyInfo(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("GetSocietyInfo")
	result, err := h.sS.GetSocietyInfo(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get society info error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(result)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to json marshal error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) UpdateSociety(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("UpdateSociety")
	err := h.sS.UpdateSociety(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to update society error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) CheckSubscriptionToPeer(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("CheckSubscriptionToPeer")
	result, err := h.fS.CheckSubscribe(r)
	if err != nil {
		logger.Error(fmt.Sprintf("check subscribe error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsn, err := json.Marshal(result)
	if err != nil {
		logger.Error(fmt.Sprintf("json marshal error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	readType := r.URL.Query().Get("type")
	var jsn []byte
	var res interface{}
	var err error
	if readType == "peer" {
		res, err = h.srS.GetUsersWithLimit(r)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to get users with limit error: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if readType == "society" {
		res, err = h.srS.GetSocietyWithLimit(r)
		if err != nil {
			logger.Error(fmt.Sprintf("check subscribe error: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	jsn, err = json.Marshal(res)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to json marshal error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetChats(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("GetChats")

	result, err := h.cS.GetChats(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get chats: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(result)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to json marshal error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) CreatePrivateChat(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("CreatePrivateChat")

	result, err := h.cS.CreatePrivateChat(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create private chat: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(result)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to json marshal error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetPrivateRecentMessages(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("GetPrivateRecentMessages")

	result, err := h.cS.GetPrivateRecentMessages(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get private recent messages: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(result)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to json marshal error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) GetAdverts(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("GetAdverts")

	result, err := h.adS.GetAdverts(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get adverts: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(result)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to json marshal: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsn)
}

func (h *Handler) CreateAdvert(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("CreateAdvert")

	result, err := h.adS.CreateAdvert(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create advert: %v", err))
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

func (h *Handler) CancelAdvert(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("CancelAdvert")

	result, err := h.adS.CancelAdvert(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to cancel advert: %v", err))
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

func (h *Handler) RestoreAdvert(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("RestoreAdvert")

	result, err := h.adS.RestoreAdvert(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to restore advert: %v", err))
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

func (h *Handler) GetOptionRequests(w http.ResponseWriter, r *http.Request) {
	logger := logger_lib.FromContext(r.Context(), config.KeyLogger)
	logger.AddFuncName("GetOptionRequests")

	result, err := h.oS.GetOptionRequests(r)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get option requests: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsn, err := json.Marshal(result)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to json marshal: %v", err))
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
		apiRouter.Post("/avatar/user", handler.SetUserAvatar)
		apiRouter.Get("/avatar/user", handler.GetAllUserAvatars)
		apiRouter.Delete("/avatar/user", handler.DeleteUserAvatar)
		apiRouter.Post("/avatar/society", handler.SetSocietyAvatar)
		apiRouter.Get("/avatar/society", handler.GetAllSocietyAvatars)
		apiRouter.Delete("/avatar/society", handler.DeleteSocietyAvatar)
		apiRouter.Get("/notification/count", handler.CountNotifications)
		apiRouter.Get("/notification", handler.GetNotifications)
		apiRouter.Patch("/notification", handler.MarkNotificationAsRead)
		apiRouter.Get("/friends/counts", handler.GetCountFriends)
		apiRouter.Get("/option/os", handler.GetOsBySearchName)
		apiRouter.Get("/option/workplace", handler.GetWorkPlaceBySearchName)
		apiRouter.Get("/option/study-place", handler.GetStudyPlaceBySearchName)
		apiRouter.Get("/option/hobby", handler.GetHobbyBySearchName)
		apiRouter.Get("/option/skill", handler.GetSkillBySearchName)
		apiRouter.Get("/option/city", handler.GetCityBySearchName)
		apiRouter.Get("/option/society-direction", handler.GetSocietyDirectionBySearchName)
		apiRouter.Post("/society", handler.CreateSociety)
		apiRouter.Get("/society", handler.GetSocietyInfo)
		apiRouter.Put("/society", handler.UpdateSociety)
		apiRouter.Post("/friends", handler.SetFriends)
		apiRouter.Delete("/friends", handler.RemoveFriends)
		apiRouter.Get("/friends/check", handler.CheckSubscriptionToPeer)
		apiRouter.Get("/peer/{uuid}", handler.PeerInfo)
		apiRouter.Get("/search", handler.Search)
		//apiRouter.Post("/society/member", handler.SubscribeToSociety)
		//apiRouter.Delete("/society/member", handler.UnsubscribeFromSociety)
		apiRouter.Get("/chat", handler.GetChats)
		apiRouter.Post("/chat/private", handler.CreatePrivateChat)
		apiRouter.Get("/chat/private/messages", handler.GetPrivateRecentMessages)
		//apiRouter.Get("/society/list", handler.GetSocietiesForUser)
		apiRouter.Get("/advert", handler.GetAdverts)
		apiRouter.Post("/advert", handler.CreateAdvert)
		apiRouter.Put("/advert/cancel", handler.CancelAdvert)
		apiRouter.Patch("/advert/restore", handler.RestoreAdvert)

		//crm routes
		apiRouter.Get("/option_requests", handler.GetOptionRequests)
	})
}
