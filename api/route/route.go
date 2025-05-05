package route

import (
	"database/sql"
	"hms-api/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *sql.DB, gin *gin.Engine) {
	publicRouter := gin.Group("");

	NewRegisterRoute(env, timeout, db, publicRouter)
	NewLoginRoute(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)
	NewDoctorRoute(env, timeout, db, publicRouter)
}