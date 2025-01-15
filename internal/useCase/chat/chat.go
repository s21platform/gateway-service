package chat

import (
	"encoding/json"
	"fmt"
	"github.com/s21platform/gateway-service/internal/config"
	"io"
	"net/http"

	chat "github.com/s21platform/chat-proto/chat-proto"
)

type Usecase struct {
	cC ChatClient
}

func New(cC ChatClient) *Usecase {
	return &Usecase{cC: cC}
}

func (u *Usecase) GetRecentMessages(r *http.Request) (*chat.GetRecentMessagesOut, error) {
	uuid := r.Context().Value(config.KeyUUID).(string)

	resp, err := u.cC.GetRecentMessages(r.Context(), &chat.GetRecentMessagesIn{Uuid: requestData.Uuid})
	if err != nil {
		return nil, fmt.Errorf("failed to u.cC.GetRecentMessages: %v", err)
	}
	return resp, nil
}
