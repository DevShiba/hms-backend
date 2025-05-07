package controller

import (
	"hms-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DoctorController struct {
	DoctorUsecase domain.DoctorUsecase
}

func NewDoctorController(usecase domain.DoctorUsecase) *DoctorController {
	return &DoctorController{
		DoctorUsecase: usecase,
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
	
	c.JSON(http.StatusNoContent, nil)
}