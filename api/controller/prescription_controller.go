package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hms-api/domain"
	"net/http"
)

type PrescriptionController struct {
	PrescriptionUsecase domain.PrescriptionRepository
}

func NewPrescriptionController(prescriptionUsecase domain.PrescriptionRepository) *PrescriptionController {
	return &PrescriptionController{
		PrescriptionUsecase: prescriptionUsecase,
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

	c.JSON(http.StatusNoContent, nil)
}
