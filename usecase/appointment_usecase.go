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
	appointment, err := pu.repository.GetAppointmentById(appointment_id)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}

func (pu *AppointmentUsecase) CreateAppointment(appointment model.Appointment) (model.Appointment, error){
	appointmentId, err := pu.repository.CreateAppointment(appointment)
	if err != nil {
		return model.Appointment{}, err
	}

	appointment.ID = appointmentId

	return appointment, nil
}

func (pu *AppointmentUsecase) UpdateAppointment(appointment_id uuid.UUID, appointment model.Appointment) (model.Appointment, error){
	appointmentId, err := pu.repository.UpdateAppointment(appointment_id, appointment)
	if err != nil {
		return model.Appointment{}, err
	}

	appointment.ID = appointmentId

	return appointment, nil
}

func (au *AppointmentUsecase) DeleteAppointment(appointment_id uuid.UUID) (*model.Appointment, error) {
	appointment, err := au.repository.DeleteAppointment(appointment_id)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}