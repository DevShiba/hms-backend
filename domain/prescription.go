package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Prescription struct {
	ID                uuid.UUID `json:"prescription_id"`
	PatientID         uuid.UUID `json:"patient_id"`
	DoctorID          uuid.UUID `json:"doctor_id"`
	MedicalRecordID   uuid.UUID `json:"medical_record_id"`
	MedicationDetails string    `json:"medication_details"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
}

type PrescriptionRepository interface {
	Create(c context.Context, prescription *Prescription) error
	Fetch(c context.Context) ([]Prescription, error)
	FetchByID(c context.Context, id uuid.UUID) (*Prescription, error)
	Update(c context.Context, prescription *Prescription) error
	Delete(c context.Context, id uuid.UUID) error
}

type PrescriptionUsecase interface {
	Create(c context.Context, prescription *Prescription) error
	Fetch(c context.Context) ([]Prescription, error)
	FetchByID(c context.Context, id uuid.UUID) (*Prescription, error)
	Update(c context.Context, prescription *Prescription) error
	Delete(c context.Context, id uuid.UUID) error
}
