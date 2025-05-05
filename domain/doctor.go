package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Doctor struct {
	ID uuid.UUID `json:"doctor_id"`
	UserId uuid.UUID `json:"user_id"`
	CRM string `json:"crm"`
	Specialty string `json:"specialty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type DoctorRepository interface {
	Create(c context.Context, doctor *Doctor) error
	Fetch(c context.Context) ([]Doctor, error)
	FetchByID(c context.Context, id uuid.UUID) (Doctor, error)
	Update(c context.Context, doctor *Doctor) error
	Delete(c context.Context, id uuid.UUID) error
}

type DoctorUsecase interface {
	Create(c context.Context, doctor *Doctor) error
	Fetch(c context.Context) ([]Doctor, error)
	FetchByID(c context.Context, id uuid.UUID) (Doctor, error)
	Update(c context.Context, doctor *Doctor) error
	Delete(c context.Context, id uuid.UUID) error
}