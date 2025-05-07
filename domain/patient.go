package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	ID uuid.UUID `json:"patient_id"`
	UserId uuid.UUID `json:"user_id"`
	CPF string `json:"cpf"`
	DateBirth string `json:"date_birth"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type PatientRepository interface {
	Create(c context.Context, patient *Patient) error
	Fetch(c context.Context) ([]Patient, error)
	FetchByID(c context.Context, id uuid.UUID) (Patient, error)
	Update(c context.Context, patient *Patient) error
	Delete(c context.Context, id uuid.UUID) error
}

type PatientUsecase interface {
	Create(c context.Context, patient *Patient) error
	Fetch(c context.Context) ([]Patient, error)
	FetchByID(c context.Context, id uuid.UUID) (Patient, error)
	Update(c context.Context, patient *Patient) error
	Delete(c context.Context, id uuid.UUID) error
}