package chat

import (
	"encoding/json"
	"fmt"
	"net/http"

	chat "github.com/s21platform/chat-proto/chat-proto"

	"github.com/s21platform/gateway-service/internal/model"
)

type Usecase struct {
	cC ChatClient
}

func New(cC ChatClient) *Usecase {
	return &Usecase{cC: cC}
}

func (u *Usecase) GetChats(r *http.Request) (*chat.GetChatsOut, error) {
	resp, err := u.cC.GetChats(r.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get chats in usecase: %v", err)
	}

	return resp, nil
}

func (u *Usecase) CreatePrivateChat(r *http.Request) (*chat.CreatePrivateChatOut, error) {
	var requestData model.PrivateChatRequest

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode request body: %v", err)
	}
	defer r.Body.Close()

	resp, err := u.cC.CreatePrivateChat(r.Context(), requestData.CompanionUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to u.cC.GetRecentMessages: %v", err)
	}

	return resp, nil
}

func (u *Usecase) GetPrivateRecentMessages(r *http.Request) (*chat.GetPrivateRecentMessagesOut, error) {
	chatUUID := r.URL.Query().Get("uuid")
	if chatUUID == "" {
		return nil, fmt.Errorf("failed to no chat UUID in request")
	}

	resp, err := u.cC.GetPrivateRecentMessages(r.Context(), chatUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get private recent messages in usecase: %v", err)
	}

	return resp, nil
}
