package usecase

import (
	"hms-api/model"
	"hms-api/repository"

	"github.com/google/uuid"
)

type AppointmentUsecase struct {
	repository repository.AppointmentRepository
}

func NewAppointmentUsecase(repo repository.AppointmentRepository) AppointmentUsecase {
	return AppointmentUsecase{
		repository: repo,
	}
}

func (pu * AppointmentUsecase) GetAppointments() ([]model.Appointment, error){
	return pu.repository.GetAppointments()
}

func (pu * AppointmentUsecase) GetAppointmentById(appointment_id uuid.UUID) (*model.Appointment, error){
	product, err := pu.repository.GetAppointmentById(appointment_id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *AppointmentUsecase) CreateAppointment(appointment model.Appointment) (model.Appointment, error){
	appointmentId, err := pu.repository.CreateAppointment(appointment)
	if err != nil {
		return model.Appointment{}, err
	}

	appointment.ID = appointmentId

	return appointment, nil
}