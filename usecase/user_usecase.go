package usecase

import (
	"hms-api/model"
	"hms-api/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return UserUsecase{
		repository: repo,
	}
}

func (uu *UserUsecase) RegisterUser(user model.User) (model.User, error){
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	user.Password = string(hashed)

	userId, err := uu.repository.RegisterUser(user)
	if err != nil {
		return model.User{}, err
	}

	user.ID = userId

	return user, nil
}

func (uu *UserUsecase) LoginUser(email, password string) (model.User, error){
	user, err := uu.repository.GetUserByEmail(email)
	if err != nil {
		return model.User{}, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return model.User{}, err
	}
	
	return user, nil
}