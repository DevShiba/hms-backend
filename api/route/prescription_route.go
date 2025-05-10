package route

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"hms-api/api/controller"
	"hms-api/bootstrap"
	"hms-api/repository"
	"hms-api/usecase"
	"time"
)

func NewPrescriptionRoute(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	pr := repository.NewPrescriptionRepository(db)
	pc := controller.NewPrescriptionController(usecase.NewPrescriptionUsecase(pr, timeout))

	group.POST("/prescriptions", pc.Create)
	group.GET("/prescriptions", pc.Fetch)
	group.GET("/prescriptions/:id", pc.FetchByID)
	group.PATCH("/prescriptions/:id", pc.Update)
	group.DELETE("/prescriptions/:id", pc.Delete)
}
