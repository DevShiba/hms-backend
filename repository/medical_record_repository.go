package repository

import (
	"context"
	"database/sql"
	"fmt"
	"hms-api/domain"

	"github.com/google/uuid"
)

type medicalRecordRepository struct {
	database *sql.DB
}

func NewMedicalRecordRepository(db *sql.DB) domain.MedicalRecordRepository {
	return &medicalRecordRepository{
		database: db,
	}
}

func (mr *medicalRecordRepository) Create(c context.Context, record *domain.MedicalRecord) error {
	query := `
        INSERT INTO medical_records (patient_id, doctor_id, diagnosis, treatment)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
    `
	return mr.database.QueryRowContext(c, query,
		record.PatientID,
		record.DoctorID,
		record.Diagnosis,
		record.Treatment,
	).Scan(&record.ID, &record.CreatedAt)
}

func (mr *medicalRecordRepository) Fetch(c context.Context) ([]domain.MedicalRecord, error) {
	query := `
        SELECT id, patient_id, doctor_id, diagnosis, treatment, created_at
        FROM medical_records
    `
	rows, err := mr.database.QueryContext(c, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching medical records: %w", err)
	}
	defer rows.Close()

	var records []domain.MedicalRecord
	for rows.Next() {
		var record domain.MedicalRecord
		if err := rows.Scan(
			&record.ID,
			&record.PatientID,
			&record.DoctorID,
			&record.Diagnosis,
			&record.Treatment,
			&record.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning medical record: %w", err)
		}
		records = append(records, record)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating medical records: %w", err)
	}

	return records, nil
}

func (mr *medicalRecordRepository) FetchByID(c context.Context, id uuid.UUID) (*domain.MedicalRecord, error) {
	query := `
        SELECT id, patient_id, doctor_id, diagnosis, treatment, created_at
        FROM medical_records
        WHERE id = $1
    `

	record := &domain.MedicalRecord{}
	err := mr.database.QueryRowContext(c, query, id).Scan(
		&record.ID,
		&record.PatientID,
		&record.DoctorID,
		&record.Diagnosis,
		&record.Treatment,
		&record.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error fetching medical record: %w", err)
	}

	return record, nil
}

func (mr *medicalRecordRepository) Update(c context.Context, record *domain.MedicalRecord) error {
	query := `
        UPDATE medical_records
        SET patient_id = $1, doctor_id = $2, diagnosis = $3, treatment = $4
        WHERE id = $5
    `

	result, err := mr.database.ExecContext(c, query,
		record.PatientID,
		record.DoctorID,
		record.Diagnosis,
		record.Treatment,
		record.ID,
	)

	if err != nil {
		return fmt.Errorf("error updating medical record: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("medical record with ID %s not found", record.ID)
	}

	return nil
}

func (mr *medicalRecordRepository) Delete(c context.Context, id uuid.UUID) error {
	query := `
        DELETE FROM medical_records
        WHERE id = $1
    `

	result, err := mr.database.ExecContext(c, query, id)
	if err != nil {
		return fmt.Errorf("error deleting medical record: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("medical record with ID %s not found", id)
	}

	return nil
}
