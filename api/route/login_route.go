package route

import (
	"database/sql"
	"hms-api/api/controller"
	"hms-api/bootstrap"
	"hms-api/internal/auditservice"
	"hms-api/repository"
	"hms-api/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoginRoute(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	alr := repository.NewAuditLogRepository(db) 
	alu := usecase.NewAuditLogUsecase(alr, timeout) 
	as := auditservice.NewService(alu)            

	lc := controller.NewLoginController( 
		usecase.NewLoginUsecase(ur, timeout),
		env,
		as, 
	)
	group.POST("/login", lc.Login)
}