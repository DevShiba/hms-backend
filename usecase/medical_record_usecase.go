package usecase

import (
	"context"
	"github.com/google/uuid"
	"hms-api/domain"
	"time"
)

type medicalRecordUsecase struct {
	medicalRecordRepository domain.MedicalRecordRepository
	contextTimeout          time.Duration
}

func NewMedicalRecordUsecase(medicalRecordRepository domain.MedicalRecordRepository, timeout time.Duration) domain.MedicalRecordUsecase {
	return &medicalRecordUsecase{
		medicalRecordRepository: medicalRecordRepository,
		contextTimeout:          timeout,
	}
}

func (mu *medicalRecordUsecase) Create(c context.Context, record *domain.MedicalRecord) error {
	ctx, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()
	return mu.medicalRecordRepository.Create(ctx, record)
}

func (mu *medicalRecordUsecase) Fetch(c context.Context) ([]domain.MedicalRecord, error) {
	ctx, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()
	return mu.medicalRecordRepository.Fetch(ctx)
}

func (mu *medicalRecordUsecase) FetchByID(c context.Context, id uuid.UUID) (*domain.MedicalRecord, error) {
	ctx, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()
	return mu.medicalRecordRepository.FetchByID(ctx, id)
}

func (mu *medicalRecordUsecase) Update(c context.Context, record *domain.MedicalRecord) error {
	ctx, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()
	return mu.medicalRecordRepository.Update(ctx, record)
}

func (mu *medicalRecordUsecase) Delete(c context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()
	return mu.medicalRecordRepository.Delete(ctx, id)
}
