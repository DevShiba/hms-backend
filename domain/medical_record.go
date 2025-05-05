package domain

import (
	"time"

	"github.com/google/uuid"
)

type MedicalRecord struct {
	Id uuid.UUID `json:"medical_record_id"`
	PatientID uuid.UUID `json:"patient_id"`
	DoctorID uuid.UUID `json:"doctor_id"`
	Diagnosis string `json:"diagnosis"`
	Treatment string `json:"treatment"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}