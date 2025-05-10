package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Action      string    `json:"action"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type AuditLogRepository interface {
	Create(c context.Context, log *AuditLog) error
	Fetch(c context.Context) ([]AuditLog, error)
	FetchByID(c context.Context, id uuid.UUID) (AuditLog, error)
	Update(c context.Context, log *AuditLog) error
	Delete(c context.Context, id uuid.UUID) error
}

type AuditLogUsecase interface {
	Create(c context.Context, log *AuditLog) error
	Fetch(c context.Context) ([]AuditLog, error)
	FetchByID(c context.Context, id uuid.UUID) (AuditLog, error)
	Update(c context.Context, log *AuditLog) error
	Delete(c context.Context, id uuid.UUID) error
}
