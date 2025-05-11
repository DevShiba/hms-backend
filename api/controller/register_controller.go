package controller

import (
	"hms-api/bootstrap"
	"hms-api/domain"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterController struct {
	RegisterUsecase domain.RegisterUsecase
	Env             *bootstrap.Env
}

var rxEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func (rc *RegisterController) Register(c *gin.Context){
	var request domain.RegisterRequest

	err := c.ShouldBind(&request)
		if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if !rxEmail.MatchString(request.Email) {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid email format"})
		return
	}

	_, err = rc.RegisterUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		ID: uuid.New(),
		Username: request.Username,
		Email: request.Email,
		Password: request.Password,
		Role: request.Role,
	}

	err = rc.RegisterUsecase.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

		accessToken, err := rc.RegisterUsecase.CreateAccessToken(&user, rc.Env.AccessTokenSecret, rc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := rc.RegisterUsecase.CreateRefreshToken(&user, rc.Env.RefreshTokenSecret, rc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := domain.RegisterResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)

}