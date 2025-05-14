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

func NewPrescriptionRoute(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	pr := repository.NewPrescriptionRepository(db)
	alr := repository.NewAuditLogRepository(db)
	alu := usecase.NewAuditLogUsecase(alr, timeout)
	as := auditservice.NewService(alu)	
	pc := controller.NewPrescriptionController(usecase.NewPrescriptionUsecase(pr, timeout), as)

	group.POST("/prescriptions", middleware.RBACMiddleware(domain.AdminRole, domain.DoctorRole), pc.Create)
	group.GET("/prescriptions", middleware.RBACMiddleware(domain.AdminRole), pc.Fetch)
	group.GET("/prescriptions/:id", middleware.RBACMiddleware(domain.AdminRole), pc.FetchByID)
	group.GET("/prescriptions/patient/:patient_id", middleware.RBACMiddleware(domain.AdminRole, domain.PatientRole), pc.FetchByPatientID)
	group.GET("/prescriptions/doctor/:doctor_id", middleware.RBACMiddleware(domain.AdminRole, domain.DoctorRole), pc.FetchByDoctorID)
	group.PATCH("/prescriptions/:id", middleware.RBACMiddleware(domain.AdminRole, domain.DoctorRole), pc.Update)
	group.DELETE("/prescriptions/:id", middleware.RBACMiddleware(domain.AdminRole), pc.Delete)
}
