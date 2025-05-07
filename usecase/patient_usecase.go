package usecase

import (
	"context"
	"hms-api/domain"
	"time"

	"github.com/google/uuid"
)

type patientUsecase struct {
	patientRepository domain.PatientRepository
	contextTimeout   time.Duration
}

func NewPatientUsecase(patientRepository domain.PatientRepository, timeout time.Duration) domain.PatientUsecase {
	return &patientUsecase{
		patientRepository: patientRepository,
		contextTimeout: timeout,
	}
}

func (pu *patientUsecase) Create(c context.Context, patient *domain.Patient) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.patientRepository.Create(ctx, patient)
}

func (pu *patientUsecase) Fetch(c context.Context) ([]domain.Patient, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.patientRepository.Fetch(ctx)
}

func (pu *patientUsecase) FetchByID(c context.Context, id uuid.UUID) (domain.Patient, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.patientRepository.FetchByID(ctx, id)
}

func (pu *patientUsecase) Update(c context.Context, patient *domain.Patient) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.patientRepository.Update(ctx, patient)
}

func (pu *patientUsecase) Delete(c context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.patientRepository.Delete(ctx, id)
}