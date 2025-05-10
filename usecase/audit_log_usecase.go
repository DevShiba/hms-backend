package usecase

import (
	"context"
	"github.com/google/uuid"
	"hms-api/domain"
	"time"
)

type auditLogUsecase struct {
	auditLogRepository domain.AuditLogRepository
	contextTimeout     time.Duration
}

func NewAuditLogUsecase(auditLogRepository domain.AuditLogRepository, timeout time.Duration) domain.AuditLogUsecase {
	return &auditLogUsecase{
		auditLogRepository: auditLogRepository,
		contextTimeout:     timeout,
	}
}

func (alu *auditLogUsecase) Create(c context.Context, auditLog *domain.AuditLog) error {
	ctx, cancel := context.WithTimeout(c, alu.contextTimeout)
	defer cancel()
	return alu.auditLogRepository.Create(ctx, auditLog)
}

func (alu *auditLogUsecase) Fetch(c context.Context) ([]domain.AuditLog, error) {
	ctx, cancel := context.WithTimeout(c, alu.contextTimeout)
	defer cancel()
	return alu.auditLogRepository.Fetch(ctx)
}

func (alu *auditLogUsecase) FetchByID(c context.Context, id uuid.UUID) (domain.AuditLog, error) {
	ctx, cancel := context.WithTimeout(c, alu.contextTimeout)
	defer cancel()
	return alu.auditLogRepository.FetchByID(ctx, id)
}

func (alu *auditLogUsecase) Update(c context.Context, auditLog *domain.AuditLog) error {
	ctx, cancel := context.WithTimeout(c, alu.contextTimeout)
	defer cancel()
	return alu.auditLogRepository.Update(ctx, auditLog)
}

func (alu *auditLogUsecase) Delete(c context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(c, alu.contextTimeout)
	defer cancel()
	return alu.auditLogRepository.Delete(ctx, id)
}
