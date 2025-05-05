package route

import (
	"database/sql"
	"hms-api/api/controller"
	"hms-api/bootstrap"
	"hms-api/repository"
	"hms-api/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRegisterRoute(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup){
	ur := repository.NewUserRepository(db)
	rc := &controller.RegisterController{
		RegisterUsecase: usecase.NewRegisterUsecase(ur, timeout),
		Env: 			    env,
	}	
	group.POST("/register", rc.Register)
}