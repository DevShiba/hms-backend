package controller

import (
	"hms-api/model"
	"hms-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type doctorController struct {
	doctorUsecase usecase.DoctorUsecase
}

func NewDoctorController(usecase usecase.DoctorUsecase) doctorController {
	return doctorController{
		doctorUsecase: usecase,
	}
}

func (d *doctorController) GetDoctors(ctx *gin.Context){
	doctors, err := d.doctorUsecase.GetDoctors()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, doctors)
}

func (d *doctorController) GetDoctorById(ctx *gin.Context){
	doctorId := ctx.Param("doctor_id")
	if doctorId == "" {
		response := model.Response{
			Message: "doctor_id is required",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	parsedUUID, err := uuid.Parse(doctorId)
	if err != nil {
		response := model.Response{
			Message: "Invalid doctor_id",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	doctor, err := d.doctorUsecase.GetDoctorById(parsedUUID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if doctor == nil {
		response := model.Response{
			Message: "Doctor not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, doctor)
}


func (d *doctorController) CreateDoctor(ctx *gin.Context){
	var doctor model.Doctor
	err := ctx.BindJSON(&doctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	createdDoctor, err := d.doctorUsecase.CreateDoctor(doctor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, createdDoctor)
}

func (d *doctorController) UpdateDoctor(ctx *gin.Context){
	doctorId := ctx.Param("doctor_id")
	if doctorId == "" {
		response := model.Response{
			Message: "doctor_id is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	parsedUUID, err := uuid.Parse(doctorId)

	if err != nil {
		response := model.Response{
			Message: "Invalid doctor_id",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var doctor model.Doctor
	err = ctx.BindJSON(&doctor);
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	updatedDoctor, err := d.doctorUsecase.UpdateDoctor(parsedUUID, doctor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedDoctor)
}

func (d *doctorController) DeleteDoctor(ctx *gin.Context){
	doctorId := ctx.Param("doctor_id")
	if doctorId == "" {
		response := model.Response{
			Message: "doctor_id is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	parsedUUID, err := uuid.Parse(doctorId)

	if err != nil {
		response := model.Response{
			Message: "Invalid doctor_id",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = d.doctorUsecase.DeleteDoctor(parsedUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}