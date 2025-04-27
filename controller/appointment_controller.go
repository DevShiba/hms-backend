package controller

import (
	"hms-api/model"
	"hms-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type appointmentController struct {
  appointmentUsecase usecase.AppointmentUsecase  
}

func NewAppointmentController(usecase usecase.AppointmentUsecase) appointmentController {
    return appointmentController{
        appointmentUsecase: usecase,  
    }
}

func (p *appointmentController) GetAppointments(ctx *gin.Context) {
    appointments, err := p.appointmentUsecase.GetAppointments()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    // Return appointments on success
    ctx.JSON(http.StatusOK, appointments)
}

func (p *appointmentController) GetAppointmentById(ctx *gin.Context){
	appointmentId := ctx.Param("appointment_id")
	if appointmentId == "" {
		response := model.Response{
			Message: "appointment_id is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	parsedUUID, err := uuid.Parse(appointmentId)
	if err != nil {
		response := model.Response {
			Message: "Invalid appointment_id",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	appointments, err := p.appointmentUsecase.GetAppointmentById(parsedUUID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if appointments == nil {
		response := model.Response{
			Message: "appointment not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, appointments)
}

func (p *appointmentController) CreateAppointment(ctx *gin.Context){
	var appointment model.Appointment
	err := ctx.BindJSON(&appointment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedAppointment, err := p.appointmentUsecase.CreateAppointment(appointment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedAppointment)
}

func (p *appointmentController) UpdateAppointment(ctx *gin.Context){
	appointmentId := ctx.Param("appointment_id")
	if appointmentId == "" {
		response := model.Response{
			Message: "appointment_id is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	parsedUUID, err := uuid.Parse(appointmentId)

	if err != nil {
		response := model.Response {
			Message: "Invalid appointment_id",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var appointment model.Appointment
	if err := ctx.BindJSON(&appointment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	updatedAppointment, err := p.appointmentUsecase.UpdateAppointment(parsedUUID, appointment)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	ctx.JSON(http.StatusOK, updatedAppointment)

}

func (a *appointmentController) DeleteAppointment(ctx *gin.Context){
	appointmentId := ctx.Param("appointment_id")
	if appointmentId == "" {
			response := model.Response{
				Message: "appointment_id is required",
			}
			ctx.JSON(http.StatusBadRequest, response)
			return
	}

	parsedUUID, err := uuid.Parse(appointmentId)
	if err != nil {
			response := model.Response {
				Message: "Invalid appointment_id",
			}
			ctx.JSON(http.StatusBadRequest, response)
			return
	}

	appointment, err := a.appointmentUsecase.DeleteAppointment(parsedUUID)
	if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
	}

	if appointment == nil {
			response := model.Response{
				Message: "appointment not found",
			}
			ctx.JSON(http.StatusNotFound, response)
			return
	}

	ctx.JSON(http.StatusOK, appointment)
}