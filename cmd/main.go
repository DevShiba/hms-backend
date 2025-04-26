package main

import (
	"hms-api/controller"
	"hms-api/db"
	"hms-api/repository"
	"hms-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	
	AppointmentRepository := repository.NewAppointmentRepository(dbConnection)
	
	AppointmentUsecase := usecase.NewAppointmentUsecase(AppointmentRepository)
	AppointmentController := controller.NewAppointmentController(AppointmentUsecase)

	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/appointments", AppointmentController.GetAppointments)
	server.POST("/appointments", AppointmentController.CreateAppointment)
	server.GET("/appointments/:appointment_id", AppointmentController.GetAppointmentById)

	server.Run(":8080")
}