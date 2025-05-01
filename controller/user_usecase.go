package controller

import (
	"hms-api/model"
	"hms-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) userController {
	return userController{
		userUsecase: usecase,
	}
}

func (u *userController) RegisterUser(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedUser, err := u.userUsecase.RegisterUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedUser)
}

func (u *userController) LoginUser(ctx *gin.Context){
	var creds struct {
		Email 	string `json:"email"`
		Password string `json:"password"`
	}
	err := ctx.BindJSON(&creds)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := u.userUsecase.LoginUser(creds.Email, creds.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}