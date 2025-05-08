package controller

import (
	"hms-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AppointmentController struct {
	AppointmentUsecase domain.AppointmentUsecase
}

func NewAppointmentController(usecase domain.AppointmentUsecase) *AppointmentController {
	return &AppointmentController{
		AppointmentUsecase: usecase,
	}
}

func (ac *AppointmentController) Create(c *gin.Context){
	var appointment domain.Appointment

	err := c.ShouldBind(&appointment)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = ac.AppointmentUsecase.Create(c, &appointment)
		if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, appointment)
}

func (ac *AppointmentController) Fetch(c *gin.Context){
	appointments, err := ac.AppointmentUsecase.Fetch(c)

		if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

func (ac *AppointmentController) FetchByID(c *gin.Context){
	appointmentID := c.Param("id")

	if appointmentID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Appointment id is required"})
		return
	}

	parsedID, err := uuid.Parse(appointmentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid appointment id"})
		return
	}

	appointment, err := ac.AppointmentUsecase.FetchByID(c, parsedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if appointment.ID == uuid.Nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Appointment not found"})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func (ac *AppointmentController) Update(c *gin.Context){
	appointmentID := c.Param("id")

	if appointmentID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Appointment id is required"})
		return
	}

	parsedID, err := uuid.Parse(appointmentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid appointment id"})
		return
	}

	var appointment domain.Appointment
	err = c.ShouldBind(&appointment)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	appointment.ID = parsedID

	err = ac.AppointmentUsecase.Update(c, &appointment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func (ac *AppointmentController) Delete(c *gin.Context){
	appointmentID := c.Param("id")

	if appointmentID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Appointment id is required"})
		return
	}

	parsedID, err := uuid.Parse(appointmentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid appointment id"})
		return
	}

	err = ac.AppointmentUsecase.Delete(c, parsedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}