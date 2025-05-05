package usecase

import (
	"context"
	"hms-api/domain"
	tokenutil "hms-api/internal"
	"time"
)

type registerUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewRegisterUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.RegisterUsecase {
	return &registerUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (ru *registerUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.userRepository.Create(ctx, user)
}

func (ru *registerUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error){
		ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
		defer cancel()
		return ru.userRepository.GetByEmail(ctx, email)
}

func (ru *registerUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error){
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (ru *registerUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error){
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}