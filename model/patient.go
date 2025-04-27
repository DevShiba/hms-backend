package model

import (
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