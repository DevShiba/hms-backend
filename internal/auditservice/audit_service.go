package auditservice

import (
	"context"
	"hms-api/domain"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	Log(ctx context.Context, userID uuid.UUID, action string, description string) error
}

type service struct {
	auditLogUsecase domain.AuditLogUsecase
}

func NewService(usecase domain.AuditLogUsecase) Service {
	return &service{
		auditLogUsecase: usecase,
	}
}

func (s *service) Log(ctx context.Context, userID uuid.UUID, action string, description string) error {
	logEntry := &domain.AuditLog{
		UserID:      userID,
		Action:      action,
		Description: description,
		CreatedAt:   time.Now(), 
	}

	return s.auditLogUsecase.Create(ctx, logEntry)
}
