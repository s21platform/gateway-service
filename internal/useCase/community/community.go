package community

import (
	"encoding/json"
	"fmt"
	"github.com/s21platform/community-service/pkg/community"
	"github.com/s21platform/gateway-service/internal/model"
	"net/http"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Usecase struct {
	cC CommunityClient
}

func New(cC CommunityClient) *Usecase {
	return &Usecase{cC: cC}
}

func (u *Usecase) SendEduLinkingCode(r *http.Request) (*emptypb.Empty, error) {
	requestData := model.SendEduLinkingCodeRequestData{}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode request body: %v", err)
	}
	defer r.Body.Close()

	resp, err := u.cC.SendEduLinkingCode(r.Context(), &requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to send code in usecase: %v", err)
	}

	return resp, nil
}

func (u *Usecase) ValidateCode(r *http.Request) (*community.ValidateCodeOut, error) {
	requestData := model.ValidateCode{}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode request body: %v", err)
	}
	defer r.Body.Close()

	resp, err := u.cC.ValidateCode(r.Context(), &requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to send code in usecase: %v", err)
	}

	return resp, nil
}
