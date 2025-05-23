package auth

import (
	"time"
)

type SignInData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"  validate:"required,min=6"`
}

type SignUpData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"  validate:"required,min=6"`
}

type RenewAccess struct {
	AccesToken string    `json:"accessToken"`
	ExpiresAt  time.Time `json:"expiresAt"`
}
