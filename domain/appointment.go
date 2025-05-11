package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type AppointmentStatus string

const (
	Scheduled AppointmentStatus = "scheduled"
	Completed AppointmentStatus = "completed"
	Canceled  AppointmentStatus = "canceled"
)

type Appointment struct {
    ID              uuid.UUID         `json:"appointment_id"`
    PatientID       uuid.UUID         `json:"patient_id"`
    DoctorID        uuid.UUID         `json:"doctor_id"`
    AppointmentDate time.Time         `json:"appointment_date"`
    Status          AppointmentStatus `json:"status"`
    Notes           string            `json:"notes"`
    CreatedAt       time.Time         `json:"created_at,omitempty"`
    UpdatedAt       time.Time         `json:"updated_at,omitempty"`
}

type AppointmentRepository interface {
    Create(c context.Context, appointment *Appointment) error
    Fetch(c context.Context) ([]Appointment, error)
    FetchByID(c context.Context, id uuid.UUID) (Appointment, error)
    FetchByPatientID(c context.Context, patientID uuid.UUID) ([]Appointment, error)
    FetchByDoctorID(c context.Context, doctorID uuid.UUID) ([]Appointment, error)
    Update(c context.Context, appointment *Appointment) error
    Delete(c context.Context, id uuid.UUID) error
}

type AppointmentUsecase interface {
    Create(c context.Context, appointment *Appointment) error
    Fetch(c context.Context) ([]Appointment, error)
    FetchByID(c context.Context, id uuid.UUID) (Appointment, error)
    FetchByPatientID(c context.Context, patientID uuid.UUID) ([]Appointment, error)
    FetchByDoctorID(c context.Context, doctorID uuid.UUID) ([]Appointment, error)
    Update(c context.Context, appointment *Appointment) error
    Delete(c context.Context, id uuid.UUID) error
}