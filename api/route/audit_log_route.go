package route

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"hms-api/api/controller"
	"hms-api/bootstrap"
	"hms-api/repository"
	"hms-api/usecase"
	"time"
)

func NewAuditLogRoute(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	alr := repository.NewAuditLogRepository(db)
	alc := controller.NewAuditLogController(usecase.NewAuditLogUsecase(alr, timeout))

	group.POST("/audit_logs", alc.Create)
	group.GET("/audit_logs", alc.Fetch)
	group.GET("/audit_logs/:id", alc.FetchByID)
	group.PATCH("/audit_logs/:id", alc.Update)
	group.DELETE("/audit_logs/:id", alc.Delete)
}
