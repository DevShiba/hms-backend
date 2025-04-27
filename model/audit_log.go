package model

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Action string    `json:"action"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}