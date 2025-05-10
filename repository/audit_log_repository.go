package repository

import (
	"context"
	"database/sql"
	"fmt"
	"hms-api/domain"

	"github.com/google/uuid"
)

type auditLogRepository struct {
	database *sql.DB
}

func NewAuditLogRepository(db *sql.DB) domain.AuditLogRepository {
	return &auditLogRepository{
		database: db,
	}
}

func (alr *auditLogRepository) Create(c context.Context, auditLog *domain.AuditLog) error {
	query := `
	INSERT INTO audit_logs (user_id, action, description, created_at)
	VALUES ($1, $2, $3, $4)
	RETURNING id
`
	err := alr.database.QueryRowContext(c, query, auditLog.UserID, auditLog.Action, auditLog.Description, auditLog.CreatedAt).Scan(&auditLog.ID)
	if err != nil {
		fmt.Println("Error executing query:", err) 
		return err
	}

	return nil
}

func (alr *auditLogRepository) Fetch(c context.Context) ([]domain.AuditLog, error) {
	query := `
	SELECT id, user_id, action, description, created_at
	FROM audit_logs
`
	rows, err := alr.database.QueryContext(c, query)
	if err != nil {
		fmt.Println("Error executing query:", err)
	}

	defer rows.Close()

	var auditLogs []domain.AuditLog

	for rows.Next() {
		var auditLog domain.AuditLog
		err = rows.Scan(
			&auditLog.ID,
			&auditLog.UserID,
			&auditLog.Action,
			&auditLog.CreatedAt,
		)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return []domain.AuditLog{}, err
		}

		auditLogs = append(auditLogs, auditLog)
	}
	return auditLogs, nil
}

func (alr *auditLogRepository) FetchByID(c context.Context, id uuid.UUID) (domain.AuditLog, error) {
	var auditLog domain.AuditLog
	query := `
	SELECT id, user_id, action, description, created_at
	FROM audit_logs 
	WHERE id = $1
`
	err := alr.database.QueryRowContext(c, query, id).Scan(
		&auditLog.ID,
		&auditLog.UserID,
		&auditLog.Action,
		&auditLog.Description,
		&auditLog.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.AuditLog{}, nil
		}
		return domain.AuditLog{}, err
	}

	return auditLog, nil
}

func (alr *auditLogRepository) Update(c context.Context, auditLog *domain.AuditLog) error {
	query := `
	UPDATE audit_logs
	SET user_id = $1, action = $2, description = $3
	WHERE id = $4
`
	_, err := alr.database.ExecContext(c, query, auditLog.UserID, auditLog.Action, auditLog.Description, auditLog.ID)

	if err != nil {
		fmt.Println("Error executing update:", err)
		return err
	}

	return nil
}

func (alr *auditLogRepository) Delete(c context.Context, id uuid.UUID) error {
	query := `
	DELETE FROM audit_logs
	WHERE id = $1
`
	_, err := alr.database.ExecContext(c, query, id)

	if err != nil {
		fmt.Println("Error executing delete:", err)
		return err
	}

	return nil
}
