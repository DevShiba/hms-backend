package usecase

import (
	"context"
	"hms-api/domain"
	"time"

	"github.com/google/uuid"
)

type doctorUsecase struct {
	doctorRepository domain.DoctorRepository
	contextTimeout time.Duration
}

func NewDoctorUsecase(doctorRepository domain.DoctorRepository, timeout time.Duration) domain.DoctorUsecase {
	return &doctorUsecase{
		doctorRepository: doctorRepository,
		contextTimeout: timeout,
	}
}

func (du *doctorUsecase) Create(c context.Context, doctor *domain.Doctor) error {
	ctx, cancel := context.WithTimeout(c, du.contextTimeout)
	defer cancel()
	return du.doctorRepository.Create(ctx, doctor)
}

func (du *doctorUsecase) Fetch(c context.Context) ([]domain.Doctor, error) {
	ctx, cancel := context.WithTimeout(c, du.contextTimeout)
	defer cancel()
	return du.doctorRepository.Fetch(ctx)
}

func (du *doctorUsecase) FetchByID(c context.Context, id uuid.UUID) (domain.Doctor, error) {
	ctx, cancel := context.WithTimeout(c, du.contextTimeout)
	defer cancel()
	return du.doctorRepository.FetchByID(ctx, id)
}

func (du *doctorUsecase) Update(c context.Context, doctor *domain.Doctor) error {
	ctx, cancel := context.WithTimeout(c, du.contextTimeout)
	defer cancel()
	return du.doctorRepository.Update(ctx, doctor)
}

func (du *doctorUsecase) Delete(c context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(c, du.contextTimeout)
	defer cancel()
	return du.doctorRepository.Delete(ctx, id)
}