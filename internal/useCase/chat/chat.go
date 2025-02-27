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

//func (u *Usecase) GetRecentMessages(r *http.Request) (*chat.GetRecentMessagesOut, error) {
//	var requestData struct {
//		Uuid string `json:"uuid"`
//	}
//
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		return nil, fmt.Errorf("failed to read request body: %v", err)
//	}
//	defer r.Body.Close()
//
//	if len(body) == 0 {
//		return nil, fmt.Errorf("request body is empty")
//	}
//
//	if err = json.Unmarshal(body, &requestData); err != nil {
//		return nil, fmt.Errorf("failed to unmarshal request body: %v", err)
//	}
//
//	resp, err := u.cC.GetRecentMessages(r.Context(), requestData.Uuid)
//	if err != nil {
//		return nil, fmt.Errorf("failed to u.cC.GetRecentMessages: %v", err)
//	}
//	return resp, nil
//}
