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

func NewPrescriptionRoute(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	pr := repository.NewPrescriptionRepository(db)
	alr := repository.NewAuditLogRepository(db)
	alu := usecase.NewAuditLogUsecase(alr, timeout)
	as := auditservice.NewService(alu)	
	pc := controller.NewPrescriptionController(usecase.NewPrescriptionUsecase(pr, timeout), as)

	group.POST("/prescriptions", pc.Create)
	group.GET("/prescriptions", pc.Fetch)
	group.GET("/prescriptions/:id", pc.FetchByID)
	group.PATCH("/prescriptions/:id", pc.Update)
	group.DELETE("/prescriptions/:id", pc.Delete)
}
