package controller

import (
	"context" // Added for audit logging
	"fmt"     // Added for audit logging descriptions
	"hms-api/domain"
	"hms-api/internal/auditservice"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DoctorController struct {
	DoctorUsecase domain.DoctorUsecase
	AuditService  auditservice.Service
}

func NewDoctorController(usecase domain.DoctorUsecase, as auditservice.Service) *DoctorController {
	return &DoctorController{
		DoctorUsecase: usecase,
		AuditService:  as,
	}
}

func (dc *DoctorController) Create(c *gin.Context) {
	var doctor domain.Doctor

	err := c.ShouldBind(&doctor)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	doctor.ID = uuid.New()

	err = dc.DoctorUsecase.Create(c, &doctor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if dc.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = dc.AuditService.Log(context.Background(), userID, "DOCTOR_CREATE", fmt.Sprintf("Doctor created with ID: %s", doctor.ID.String()))
		}()
	}

	c.JSON(http.StatusCreated, doctor)
}

func (dc *DoctorController) Fetch(c *gin.Context) {
	doctors, err := dc.DoctorUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, doctors)
}

func (dc *DoctorController) FetchByID(c *gin.Context) {
	doctorID := c.Param("id")
	if doctorID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "doctor id is required"})
		return
	}

	parsedID, err := uuid.Parse(doctorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid doctor id format"})
		return
	}

	doctor, err := dc.DoctorUsecase.FetchByID(c, parsedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if doctor.ID == uuid.Nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Doctor not found"})
		return
	}

	if dc.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = dc.AuditService.Log(context.Background(), userID, "DOCTOR_FETCH_BY_ID", fmt.Sprintf("Fetched doctor with ID: %s", doctor.ID.String()))
		}()
	}

	c.JSON(http.StatusOK, doctor)
}

func (dc *DoctorController) Update(c *gin.Context) {
	doctorID := c.Param("id")
	if doctorID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "doctor id is required"})
		return
	}

	parsedID, err := uuid.Parse(doctorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid doctor id format"})
		return
	}

	var doctor domain.Doctor
	err = c.ShouldBind(&doctor)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	doctor.ID = parsedID

	err = dc.DoctorUsecase.Update(c, &doctor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Audit Log
	if dc.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = dc.AuditService.Log(context.Background(), userID, "DOCTOR_UPDATE", fmt.Sprintf("Updated doctor with ID: %s", doctor.ID.String()))
		}()
	}

	c.JSON(http.StatusOK, doctor)
}

func (dc *DoctorController) Delete(c *gin.Context) {
	doctorID := c.Param("id")
	if doctorID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "doctor id is required"})
		return
	}

	parsedID, err := uuid.Parse(doctorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid doctor id format"})
		return
	}

	err = dc.DoctorUsecase.Delete(c, parsedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Audit Log
	if dc.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = dc.AuditService.Log(context.Background(), userID, "DOCTOR_DELETE", fmt.Sprintf("Deleted doctor with ID: %s", parsedID.String()))
		}()
	}

	c.JSON(http.StatusNoContent, nil)
}