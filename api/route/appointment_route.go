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

func NewAppointmentRoute(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup){
	ar := repository.NewAppointmentRepository(db)
	alr := repository.NewAuditLogRepository(db)
	alu := usecase.NewAuditLogUsecase(alr, timeout)
	as := auditservice.NewService(alu)
	ac := controller.NewAppointmentController(usecase.NewAppointmentUsecase(ar, timeout), as)

	group.POST("/appointments", ac.Create)
	group.GET("/appointments", ac.Fetch)
	group.GET("/appointments/:id", ac.FetchByID)
	group.GET("/appointments/patient/:patient_id", ac.FetchByPatientID)
	group.GET("/appointments/doctor/:doctor_id", ac.FetchByDoctorID)
	group.PATCH("/appointments/:id", ac.Update)
	group.DELETE("/appointments/:id", ac.Delete)
}