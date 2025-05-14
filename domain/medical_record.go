package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type MedicalRecord struct {
	ID        uuid.UUID `json:"medical_record_id"`
	PatientID uuid.UUID `json:"patient_id"`
	DoctorID  uuid.UUID `json:"doctor_id"`
	Diagnosis string    `json:"diagnosis"`
	Treatment string    `json:"treatment"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type MedicalRecordRepository interface {
	Create(c context.Context, record *MedicalRecord) error
	Fetch(c context.Context) ([]MedicalRecord, error)
	FetchByID(c context.Context, id uuid.UUID) (*MedicalRecord, error)
	FetchByDoctorID(c context.Context, doctorID uuid.UUID) ([]MedicalRecord, error)
	Update(c context.Context, record *MedicalRecord) error
	Delete(c context.Context, id uuid.UUID) error
}

type MedicalRecordUsecase interface {
	Create(c context.Context, record *MedicalRecord) error
	Fetch(c context.Context) ([]MedicalRecord, error)
	FetchByID(c context.Context, id uuid.UUID) (*MedicalRecord, error)
	FetchByDoctorID(c context.Context, doctorID uuid.UUID) ([]MedicalRecord, error) 
	Update(c context.Context, record *MedicalRecord) error
	Delete(c context.Context, id uuid.UUID) error
}
