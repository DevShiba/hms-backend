package domain

import "context"

type RegisterRequest struct {
	Username  string    `json:"username"`
	Email 	  string    `json:"email"`
	Password  string    `json:"password"`
	Role      UserRole  `json:"role"`
}

type RegisterResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RegisterUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}