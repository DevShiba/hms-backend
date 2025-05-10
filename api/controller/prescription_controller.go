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

type PrescriptionController struct {
	PrescriptionUsecase domain.PrescriptionRepository
	AuditService  auditservice.Service
}

func NewPrescriptionController(prescriptionUsecase domain.PrescriptionRepository, as auditservice.Service) *PrescriptionController {
	return &PrescriptionController{
		PrescriptionUsecase: prescriptionUsecase,
		AuditService: as,
	}
}

func (pc *PrescriptionController) Create(c *gin.Context) {
	var prescription domain.Prescription

	err := c.ShouldBind(&prescription)
	if err != nil {
		c.JSON(400, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = pc.PrescriptionUsecase.Create(c, &prescription)
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
			_ = pc.AuditService.Log(context.Background(), userID, "PRESCRIPTION_CREATE", fmt.Sprintf("Doctor created with ID: %s", prescription.ID.String()))
		}()
	}

	c.JSON(http.StatusCreated, prescription)
}

func (pc *PrescriptionController) Fetch(c *gin.Context) {
	prescriptions, err := pc.PrescriptionUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	c.JSON(http.StatusOK, prescriptions)
}

func (pc *PrescriptionController) FetchByID(c *gin.Context) {
	prescriptionID := c.Param("id")
	if prescriptionID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "prescription id is required"})
		return
	}

	parsedID, err := uuid.Parse(prescriptionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid prescription id format"})
		return
	}

	prescription, err := pc.PrescriptionUsecase.FetchByID(c, parsedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if prescription.ID == uuid.Nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Prescription not found"})
		return
	}

	if pc.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = pc.AuditService.Log(context.Background(), userID, "PRESCRIPTION_FETCH_BY_ID", fmt.Sprintf("Prescription fetched with ID: %s", prescription.ID.String()))
		}()
	}

	c.JSON(http.StatusOK, prescription)
}

func (pc *PrescriptionController) Update(c *gin.Context) {
	prescriptionID := c.Param("id")
	if prescriptionID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "prescription id is required"})
		return
	}

	parsedID, err := uuid.Parse(prescriptionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid prescription id format"})
		return
	}

	var prescription domain.Prescription
	err = c.ShouldBind(&prescription)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	prescription.ID = parsedID

	err = pc.PrescriptionUsecase.Update(c, &prescription)
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
			_ = pc.AuditService.Log(context.Background(), userID, "PRESCRIPTION_UPDATE", fmt.Sprintf("Prescription updated with ID: %s", prescription.ID.String()))
		}()
	}

	c.JSON(http.StatusOK, prescription)
}

func (pc *PrescriptionController) Delete(c *gin.Context) {
	prescriptionID := c.Param("id")
	if prescriptionID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "prescription id is required"})
		return
	}

	parsedID, err := uuid.Parse(prescriptionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid prescription id format"})
		return
	}

	err = pc.PrescriptionUsecase.Delete(c, parsedID)
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
			_ = pc.AuditService.Log(context.Background(), userID, "PRESCRIPTION_DELETE", fmt.Sprintf("Prescription deleted with ID: %s", parsedID.String()))
		}()
	}

	c.JSON(http.StatusNoContent, nil)
}
