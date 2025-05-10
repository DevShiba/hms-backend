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

func NewDoctorRoute(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	dr := repository.NewDoctorRepository(db)
	alr := repository.NewAuditLogRepository(db) 
	alu := usecase.NewAuditLogUsecase(alr, timeout) 
	as := auditservice.NewService(alu)        
	dc := controller.NewDoctorController(usecase.NewDoctorUsecase(dr, timeout), as)

	group.POST("/doctors", dc.Create)
	group.GET("/doctors", dc.Fetch)
	group.GET("/doctors/:id", dc.FetchByID)
	group.PATCH("/doctors/:id", dc.Update)
	group.DELETE("/doctors/:id", dc.Delete)
}