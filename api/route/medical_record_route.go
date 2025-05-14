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

func NewMedicalRecordRoute(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	mrr := repository.NewMedicalRecordRepository(db)
	alr := repository.NewAuditLogRepository(db)
	alu := usecase.NewAuditLogUsecase(alr, timeout)
	as := auditservice.NewService(alu)
	mrc := controller.NewMedicalRecordController(usecase.NewMedicalRecordUsecase(mrr, timeout), as)

	group.POST("/medical_records", middleware.RBACMiddleware(domain.AdminRole, domain.DoctorRole), mrc.Create)
	group.GET("/medical_records", middleware.RBACMiddleware(domain.AdminRole), mrc.Fetch)
	group.GET("/medical_records/:id", middleware.RBACMiddleware(), mrc.FetchByID)
	group.GET("/medical_records/doctor/:doctor_id", middleware.RBACMiddleware(domain.AdminRole, domain.DoctorRole), mrc.FetchByDoctorID)
	group.PATCH("/medical_records/:id", middleware.RBACMiddleware(domain.AdminRole, domain.DoctorRole), mrc.Update)
	group.DELETE("/medical_records/:id", middleware.RBACMiddleware(domain.AdminRole), mrc.Delete)
}
