package route

import (
	"database/sql"
	"hms-api/api/controller"
	"hms-api/api/middleware"
	"hms-api/bootstrap"
	"hms-api/domain"
	"hms-api/repository"
	"hms-api/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewAuditLogRoute(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	alr := repository.NewAuditLogRepository(db)
	alc := controller.NewAuditLogController(usecase.NewAuditLogUsecase(alr, timeout))

	group.POST("/audit_logs", alc.Create)
	group.GET("/audit_logs",  middleware.RBACMiddleware(domain.AdminRole), alc.Fetch)
	group.GET("/audit_logs/:id", middleware.RBACMiddleware(domain.AdminRole), alc.FetchByID)
	group.PATCH("/audit_logs/:id",  middleware.RBACMiddleware(domain.AdminRole), alc.Update)
	group.DELETE("/audit_logs/:id",  middleware.RBACMiddleware(domain.AdminRole), alc.Delete)
}
