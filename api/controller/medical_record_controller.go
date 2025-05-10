package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hms-api/domain"
	"net/http"
)

type MedicalRecordController struct {
	MedicalRecordUsecase domain.MedicalRecordUsecase
}

func NewMedicalRecordController(medicalRecordUsecase domain.MedicalRecordUsecase) *MedicalRecordController {
	return &MedicalRecordController{
		MedicalRecordUsecase: medicalRecordUsecase,
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

	c.JSON(http.StatusOK, record)
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

	c.JSON(http.StatusNoContent, nil)
}
