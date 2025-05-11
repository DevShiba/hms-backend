package controller

import (
	"context"
	"fmt"
	"hms-api/domain"
	"hms-api/internal/auditservice"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AppointmentController struct {
	AppointmentUsecase domain.AppointmentUsecase
	AuditService  auditservice.Service
}

func NewAppointmentController(usecase domain.AppointmentUsecase, as auditservice.Service) *AppointmentController {
	return &AppointmentController{
		AppointmentUsecase: usecase,
		AuditService: as,
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

	if ac.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = ac.AuditService.Log(context.Background(), userID, "APPOINTMENT_CREATE", fmt.Sprintf("Appointment created with ID: %s", appointment.ID.String()))
		}()
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

	if ac.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = ac.AuditService.Log(context.Background(), userID, "DOCTOR_CREATE", fmt.Sprintf("Appointment fetched with ID: %s", appointmentID))
		}()
	}

	c.JSON(http.StatusOK, appointment)
}

func (ac *AppointmentController) FetchByPatientID(c *gin.Context){
	patientID := c.Param("patient_id")
	if patientID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Patient id is required"})
		return
	}

	parsedID, err := uuid.Parse(patientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid patient id"})
		return
	}

	appointments, err := ac.AppointmentUsecase.FetchByPatientID(c, parsedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if appointments == nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "No appointments found for this patient"})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

func (ac *AppointmentController) FetchByDoctorID(c *gin.Context){
	doctorID := c.Param("doctor_id")
	if doctorID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Doctor id is required"})
		return
	}

	parsedID, err := uuid.Parse(doctorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid doctor id"})
		return
	}

	appointments, err := ac.AppointmentUsecase.FetchByDoctorID(c, parsedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if appointments == nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "No appointments found for this doctor"})
		return
	}

	c.JSON(http.StatusOK, appointments)
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

	if ac.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = ac.AuditService.Log(context.Background(), userID, "APPOINTMENT_UPDATE", fmt.Sprintf("Appointment updated with ID: %s", appointment.ID.String()))
		}()
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

	if ac.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = ac.AuditService.Log(context.Background(), userID, "APPOINTMENT_DELETE", fmt.Sprintf("Appointment deleted with ID: %s", appointmentID))
		}()
	}

	c.JSON(http.StatusNoContent, nil)
}