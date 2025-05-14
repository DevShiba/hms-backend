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

func NewAppointmentRoute(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup){
	ar := repository.NewAppointmentRepository(db)
	alr := repository.NewAuditLogRepository(db)
	alu := usecase.NewAuditLogUsecase(alr, timeout)
	as := auditservice.NewService(alu)
	ac := controller.NewAppointmentController(usecase.NewAppointmentUsecase(ar, timeout), as)

	group.POST("/appointments", middleware.RBACMiddleware(domain.AdminRole, domain.DoctorRole, domain.PatientRole),ac.Create)
	group.GET("/appointments", middleware.RBACMiddleware(domain.AdminRole), ac.Fetch)
	group.GET("/appointments/:id", middleware.RBACMiddleware(domain.AdminRole),ac.FetchByID)
	group.GET("/appointments/patient/:patient_id",  middleware.RBACMiddleware(domain.AdminRole, domain.PatientRole),ac.FetchByPatientID)
	group.GET("/appointments/doctor/:doctor_id", middleware.RBACMiddleware(domain.AdminRole, domain.DoctorRole), ac.FetchByDoctorID)
	group.PATCH("/appointments/:id", middleware.RBACMiddleware(domain.AdminRole, domain.DoctorRole), ac.Update)
	group.DELETE("/appointments/:id", middleware.RBACMiddleware(domain.AdminRole, domain.DoctorRole), ac.Delete)
}