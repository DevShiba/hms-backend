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

func NewPatientRoute(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup){
	pr := repository.NewPatientRepository(db)
	alr := repository.NewAuditLogRepository(db)
	alu := usecase.NewAuditLogUsecase(alr, timeout)
	as := auditservice.NewService(alu)
	pc := controller.NewPatientController(usecase.NewPatientUsecase(pr, timeout), as)

	group.POST("/patients", pc.Create)
	group.GET("/patients", pc.Fetch)
	group.GET("/patients/:id", pc.FetchByID)
	group.GET("/patients/doctor/:doctor_id", pc.FetchByDoctorID)
	group.PATCH("/patients/:id", pc.Update)
	group.DELETE("/patients/:id", pc.Delete)
}