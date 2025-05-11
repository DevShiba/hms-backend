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

type MedicalRecordController struct {
	MedicalRecordUsecase domain.MedicalRecordUsecase
	AuditService  auditservice.Service
}

func NewMedicalRecordController(medicalRecordUsecase domain.MedicalRecordUsecase, as auditservice.Service) *MedicalRecordController {
	return &MedicalRecordController{
		MedicalRecordUsecase: medicalRecordUsecase,
		AuditService: as,
	}
}

func (mrc *MedicalRecordController) Create(c *gin.Context) {
	var record domain.MedicalRecord

	err := c.ShouldBind(&record)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = mrc.MedicalRecordUsecase.Create(c, &record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if mrc.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = mrc.AuditService.Log(context.Background(), userID, "MEDICAL_RECORD_CREATE", fmt.Sprintf("Medical Record created with ID: %s", record.ID.String()))
		}()
	}

	c.JSON(http.StatusCreated, record)
}

func (mrc *MedicalRecordController) Fetch(c *gin.Context) {
	records, err := mrc.MedicalRecordUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

func (mrc *MedicalRecordController) FetchByID(c *gin.Context) {
	recordID := c.Param("id")
	if recordID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "record id is required"})
		return
	}

	parsedID, err := uuid.Parse(recordID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid record id format"})
		return
	}

	record, err := mrc.MedicalRecordUsecase.FetchByID(c, parsedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if record.ID == uuid.Nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Record not found"})
		return
	}

	if mrc.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = mrc.AuditService.Log(context.Background(), userID, "MEDICAL_RECORD_FETCH_BY_ID", fmt.Sprintf("Medical Record fetched with ID: %s", parsedID.String()))
		}()
	}

	c.JSON(http.StatusOK, record)
}

func (mrc *MedicalRecordController) FetchByDoctorID(c *gin.Context) {
	doctorID := c.Param("doctor_id")
	if doctorID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "doctor id is required"})
		return
	}

	parsedID, err := uuid.Parse(doctorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid doctor id format"})
		return
	}

	records, err := mrc.MedicalRecordUsecase.FetchByDoctorID(c, parsedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if records == nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "No records found for this doctor"})
		return
	}
	
	c.JSON(http.StatusOK, records)
}

func (mrc *MedicalRecordController) Update(c *gin.Context) {
	recordID := c.Param("id")
	if recordID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "record id is required"})
	}

	parsedID, err := uuid.Parse(recordID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid record id format"})
		return
	}

	var record domain.MedicalRecord
	err = c.ShouldBind(&record)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	record.ID = parsedID

	err = mrc.MedicalRecordUsecase.Update(c, &record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if mrc.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = mrc.AuditService.Log(context.Background(), userID, "MEDICAL_RECORD_UPDATE", fmt.Sprintf("Medical Record updated with ID: %s", record.ID.String()))
		}()
	}

	c.JSON(http.StatusOK, record)
}

func (mrc *MedicalRecordController) Delete(c *gin.Context) {
	recordID := c.Param("id")
	if recordID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "record id is required"})
		return
	}

	parsedID, err := uuid.Parse(recordID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid record id format"})
		return
	}

	err = mrc.MedicalRecordUsecase.Delete(c, parsedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if mrc.AuditService != nil {
		userIDCtx, _ := c.Get("x-user-id")
		userID := uuid.Nil
		if id, ok := userIDCtx.(uuid.UUID); ok {
			userID = id
		}
		go func() {
			_ = mrc.AuditService.Log(context.Background(), userID, "MEDICAL_RECORD_DELETE", fmt.Sprintf("Medical Record deleted with ID: %s", parsedID.String()))
		}()
	}

	c.JSON(http.StatusNoContent, nil)
}
