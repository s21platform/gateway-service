package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	authproto "github.com/s21platform/auth-service/pkg/auth"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/s21platform/gateway-service/internal/model"
	"github.com/s21platform/gateway-service/internal/rpc/auth"
)

type Usecase struct {
	aC AuthClient
}

func New(aC AuthClient) *Usecase {
	return &Usecase{aC: aC}
}

func (uc *Usecase) Login(ctx context.Context, username string, password string) (*auth.JWT, error) {
	return uc.aC.DoLogin(ctx, username, password)
}

func (uc *Usecase) CheckEmailAvailability(r *http.Request) (*model.EmailResponse, error) {
	email := r.URL.Query().Get("email")
	if email == "" {
		return nil, fmt.Errorf("failed to no email found in request")
	}

	res, err := uc.aC.CheckEmailAvailability(r.Context(), email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email in usecase: %v", err)
	}

	return &model.EmailResponse{IsAvailable: res.IsAvailable}, nil
}

func (uc *Usecase) SendUserVerificationCode(r *http.Request) (*authproto.SendUserVerificationCodeOut, error) {
	var requestData model.CodeRequest

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode request body: %v", err)
	}
	defer r.Body.Close()

	resp, err := uc.aC.SendUserVerificationCode(r.Context(), requestData.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to send code in usecase: %v", err)
	}

	return resp, nil
}

func (uc *Usecase) RegisterUser(r *http.Request) (*emptypb.Empty, error) {
	var requestData model.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode request body: %v", err)
	}
	defer r.Body.Close()

	resp, err := uc.aC.RegisterUser(r.Context(), &requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to register in usecase: %v", err)
	}

	return resp, nil
}

func (uc *Usecase) LoginV2(r *http.Request) (*authproto.LoginV2Out, error) {
	var requestData model.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode request body: %v", err)
	}
	defer r.Body.Close()

	resp, err := uc.aC.LoginV2(r.Context(), requestData.Login, requestData.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to login in usecase: %v", err)
	}

	return resp, nil
}
