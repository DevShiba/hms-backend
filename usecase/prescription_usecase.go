package usecase

import (
	"context"
	"hms-api/domain"
	"time"

	"github.com/google/uuid"
)

type prescriptionUsecase struct {
	prescriptionRepository domain.PrescriptionRepository
	contextTimeout         time.Duration
}

func NewPrescriptionUsecase(prescriptionRepository domain.PrescriptionRepository, timeout time.Duration) domain.PrescriptionUsecase {
	return &prescriptionUsecase{
		prescriptionRepository: prescriptionRepository,
		contextTimeout:         timeout,
	}
}

func (pu *prescriptionUsecase) Create(c context.Context, prescription *domain.Prescription) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.prescriptionRepository.Create(ctx, prescription)
}

func (pu *prescriptionUsecase) Fetch(c context.Context) ([]domain.Prescription, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.prescriptionRepository.Fetch(ctx)
}

func (pu *prescriptionUsecase) FetchByID(c context.Context, id uuid.UUID) (*domain.Prescription, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.prescriptionRepository.FetchByID(ctx, id)
}

func (pu *prescriptionUsecase) FetchByPatientID(c context.Context, patientID uuid.UUID) ([]domain.Prescription, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.prescriptionRepository.FetchByPatientID(ctx, patientID)
}

func (pu *prescriptionUsecase) FetchByDoctorID(c context.Context, doctorID uuid.UUID) ([]domain.Prescription, error){
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.prescriptionRepository.FetchByDoctorID(ctx, doctorID)
}

func (pu *prescriptionUsecase) Update(c context.Context, prescription *domain.Prescription) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.prescriptionRepository.Update(ctx, prescription)
}

func (pu *prescriptionUsecase) Delete(c context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.prescriptionRepository.Delete(ctx, id)
}
