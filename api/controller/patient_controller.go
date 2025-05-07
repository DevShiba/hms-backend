package controller

import (
	"hms-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PatientController struct {
	PatientUsecase domain.PatientUsecase
}

func NewPatientController(usecase domain.PatientUsecase) *PatientController {
	return &PatientController{
		PatientUsecase: usecase,
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

	c.JSON(http.StatusNoContent, nil)
}