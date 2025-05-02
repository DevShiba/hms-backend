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

	UserRepository := repository.NewUserRepository(dbConnection)
	UserUsecase := usecase.NewUserUsecase(UserRepository)
	UserController := controller.NewUserController(UserUsecase)

	DoctorRepository := repository.NewDoctorRepository(dbConnection)
	DoctorUsecase := usecase.NewDoctorUsecase(DoctorRepository)
	DoctorController := controller.NewDoctorController(DoctorUsecase)

	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/doctors", DoctorController.GetDoctors)
	server.GET("/doctors/:doctor_id", DoctorController.GetDoctorById)
	server.POST("/doctors", DoctorController.CreateDoctor)
	server.PATCH("/doctors/:doctor_id", DoctorController.UpdateDoctor)
	server.DELETE("/doctors/:doctor_id", DoctorController.DeleteDoctor)
	server.GET("/appointments", AppointmentController.GetAppointments)
	server.POST("/register", UserController.RegisterUser)
	server.POST("/login", UserController.LoginUser)
	server.POST("/appointments", AppointmentController.CreateAppointment)
	server.GET("/appointments/:appointment_id", AppointmentController.GetAppointmentById)
	server.PATCH("/appointments/:appointment_id", AppointmentController.UpdateAppointment)
	server.DELETE("/appointments/:appointment_id", AppointmentController.DeleteAppointment)


	server.Run(":8080")
}