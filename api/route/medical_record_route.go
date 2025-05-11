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

func NewMedicalRecordRoute(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	mrr := repository.NewMedicalRecordRepository(db)
	alr := repository.NewAuditLogRepository(db)
	alu := usecase.NewAuditLogUsecase(alr, timeout)
	as := auditservice.NewService(alu)
	mrc := controller.NewMedicalRecordController(usecase.NewMedicalRecordUsecase(mrr, timeout), as)

	group.POST("/medical_records", mrc.Create)
	group.GET("/medical_records", mrc.Fetch)
	group.GET("/medical_records/:id", mrc.FetchByID)
	group.GET("/medical_records/doctor/:doctor_id", mrc.FetchByDoctorID)
	group.PATCH("/medical_records/:id", mrc.Update)
	group.DELETE("/medical_records/:id", mrc.Delete)
}
