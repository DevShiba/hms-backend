package usecase

import (
	"context"
	"hms-api/domain"
	"time"

	"github.com/google/uuid"
)

type appointmentUsecase struct {
	appointmentRepository domain.AppointmentRepository
	contextTimeout time.Duration
}

func NewAppointmentUsecase(appointmentRepository domain.AppointmentRepository, timeout time.Duration) domain.AppointmentUsecase {
	return &appointmentUsecase{
		appointmentRepository: appointmentRepository,
		contextTimeout: timeout,
	}
}

func (au *appointmentUsecase) Create(c context.Context, appointment *domain.Appointment) error {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	return au.appointmentRepository.Create(ctx, appointment)
}

func (au *appointmentUsecase) Fetch(c context.Context) ([]domain.Appointment, error){
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	return au.appointmentRepository.Fetch(ctx)
}

func (au *appointmentUsecase) FetchByID(c context.Context, id uuid.UUID) (domain.Appointment, error){
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	return au.appointmentRepository.FetchByID(ctx, id)
}

func (au *appointmentUsecase) Update(c context.Context, appointment *domain.Appointment) error {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	return au.appointmentRepository.Update(ctx, appointment)
}

func (au *appointmentUsecase) Delete(c context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	return au.appointmentRepository.Delete(ctx, id)
}