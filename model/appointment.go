package model

import (
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
    ID             uuid.UUID         `json:"appointment_id"`
    AppointmentDate time.Time        `json:"appointment_date"`
    Status         AppointmentStatus `json:"status"`
    Notes          string            `json:"notes"`
    CreatedAt      time.Time         `json:"created_at,omitempty"`
    UpdatedAt      time.Time         `json:"updated_at,omitempty"`
}