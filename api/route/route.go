package route

import (
	"database/sql"
	"hms-api/api/middleware"
	"hms-api/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *sql.DB, gin *gin.Engine) {
	publicRouter := gin.Group("");

	NewRegisterRoute(env, timeout, db, publicRouter)
	NewLoginRoute(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")

	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	NewDoctorRoute(env, timeout, db, protectedRouter)
	NewPatientRoute(env, timeout, db, protectedRouter)
	NewAppointmentRoute(env, timeout, db, protectedRouter)
}