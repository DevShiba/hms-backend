package model

import (
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