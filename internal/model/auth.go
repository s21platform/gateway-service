package model

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UUID        string `json:"uid"`
	Username    string `json:"username"`
	Role        string `json:"role"`
	AccessToken string `json:"accessToken"`
	Exp         int64  `json:"exp"`
	jwt.RegisteredClaims
}

type CheckAuth struct {
	IsAuth bool   `json:"isAuthenticated"`
	Error  string `json:"error,omitempty"`
}

type EmailResponse struct {
	IsAvailable bool `json:"is_available"`
}

type CodeRequest struct {
	Email string `json:"email"`
}

type RegisterRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Code            string `json:"code"`
	CodeLookupUUID  string `json:"code_lookup_uuid"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginV2Response struct {
	AccessToken string `json:"access_token"`
}
