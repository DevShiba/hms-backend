package route

import (
	"database/sql"
	"hms-api/api/controller"
	"hms-api/api/middleware"
	"hms-api/bootstrap"
	"hms-api/domain"
	"hms-api/internal/auditservice"
	"hms-api/repository"
	"hms-api/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewPatientRoute(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup){
	pr := repository.NewPatientRepository(db)
	alr := repository.NewAuditLogRepository(db)
	alu := usecase.NewAuditLogUsecase(alr, timeout)
	as := auditservice.NewService(alu)
	pc := controller.NewPatientController(usecase.NewPatientUsecase(pr, timeout), as)

	group.POST("/patients", middleware.RBACMiddleware(domain.AdminRole, domain.DoctorRole), pc.Create)
	group.GET("/patients", middleware.RBACMiddleware(domain.AdminRole), pc.Fetch)
	group.GET("/patients/:id", middleware.RBACMiddleware(domain.AdminRole), pc.FetchByID)
	group.GET("/patients/doctor/:doctor_id", middleware.RBACMiddleware(domain.AdminRole, domain.DoctorRole), pc.FetchByDoctorID)
	group.PATCH("/patients/:id", middleware.RBACMiddleware(domain.AdminRole, domain.DoctorRole), pc.Update)
	group.DELETE("/patients/:id", middleware.RBACMiddleware(domain.AdminRole), pc.Delete)
}