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

func NewMedicalRecordRoute(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	mrr := repository.NewMedicalRecordRepository(db)
	mrc := controller.NewMedicalRecordController(usecase.NewMedicalRecordUsecase(mrr, timeout))

	group.POST("/medical_records", mrc.Create)
	group.GET("/medical_records", mrc.Fetch)
	group.GET("/medical_records/:id", mrc.FetchByID)
	group.PATCH("/medical_records/:id", mrc.Update)
	group.DELETE("/medical_records/:id", mrc.Delete)
}
