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

type PatientController struct {
	PatientUsecase domain.PatientUsecase
	AuditService  auditservice.Service
}

func NewPatientController(usecase domain.PatientUsecase, as auditservice.Service) *PatientController {
	return &PatientController{
		PatientUsecase: usecase,
		AuditService:  as,
	}
}

func (pc *PatientController) Create(c *gin.Context){
	var patient domain.Patient

	err := c.ShouldBind(&patient)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	patient.ID = uuid.New()

	err = pc.PatientUsecase.Create(c, &patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if pc.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = pc.AuditService.Log(context.Background(), userID, "PATIENT_CREATE", fmt.Sprintf("Patient created with ID: %s", patient.ID.String()))
		}()
	}

	c.JSON(http.StatusCreated, patient)
}

func (pc *PatientController) Fetch(c *gin.Context){
	patients, err := pc.PatientUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, patients)
}

func (pc *PatientController) FetchByID(c *gin.Context){
	patientID := c.Param("id")
	if patientID == ""{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "patient id is required"})
		return
	}

	parsedID, err := uuid.Parse(patientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "invalid patient id"})
		return
	}

	patient, err := pc.PatientUsecase.FetchByID(c, parsedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if patient.ID == uuid.Nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Doctor not found"})
		return
	}

	if pc.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = pc.AuditService.Log(context.Background(), userID, "PATIENT_FETCH_BY_ID", fmt.Sprintf("Patient fetched with ID: %s", patient.ID.String()))
		}()
	}

	c.JSON(http.StatusOK, patient)
}

func (pc *PatientController) Update(c *gin.Context){
	patientID := c.Param("id")

	if patientID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "patient id is required"})
		return
	}

	parsedID, err := uuid.Parse(patientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "invalid patient id"})
		return
	}

	var patient domain.Patient
	err = c.ShouldBind(&patient)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	patient.ID = parsedID

	err = pc.PatientUsecase.Update(c, &patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if pc.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = pc.AuditService.Log(context.Background(), userID, "PATIENT_UPDATE", fmt.Sprintf("Patient updated with ID: %s", patient.ID.String()))
		}()
	}

	c.JSON(http.StatusOK, patient)
}

func (pc *PatientController) Delete(c *gin.Context){
	patientID := c.Param("id")

	if patientID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "patient id is required"})
		return
	}

	parsedID, err := uuid.Parse(patientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "invalid patient id"})
		return
	}

	err = pc.PatientUsecase.Delete(c, parsedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if pc.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = pc.AuditService.Log(context.Background(), userID, "PATIENT_DELETE", fmt.Sprintf("Patient deleted with ID: %s", parsedID.String()))
		}()
	}

	c.JSON(http.StatusNoContent, nil)
}