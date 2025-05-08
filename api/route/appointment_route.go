package route

import (
	"database/sql"
	"hms-api/api/controller"
	"hms-api/bootstrap"
	"hms-api/repository"
	"hms-api/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewAppointmentRoute(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup){
	ar := repository.NewAppointmentRepository(db)
	ac := controller.NewAppointmentController(usecase.NewAppointmentUsecase(ar, timeout))

	group.POST("/appointments", ac.Create)
	group.GET("/appointments", ac.Fetch)
	group.GET("/appointments/:id", ac.FetchByID)
	group.PATCH("/appointments/:id", ac.Update)
	group.DELETE("/appointments/:id", ac.Delete)
}