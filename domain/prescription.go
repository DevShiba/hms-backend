package domain

import (
	"time"

	"github.com/google/uuid"
)

type Prescription struct {
	ID uuid.UUID `json:"prescription_id"`
	PatientID uuid.UUID `json:"patient_id"`
	DoctorID uuid.UUID `json:"doctor_id"`
	MedicalRecordID uuid.UUID `json:"medical_record_id"`
	MedicationDetails string `json:"medication_details"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}