package repository

import (
	"context"
	"database/sql"
	"fmt"
	"hms-api/domain"

	"github.com/google/uuid"
)

type prescriptionRepository struct {
	database *sql.DB
}

func NewPrescriptionRepository(db *sql.DB) domain.PrescriptionRepository {
	return &prescriptionRepository{
		database: db,
	}
}

func (pr *prescriptionRepository) Create(c context.Context, prescription *domain.Prescription) error {
	query := `
		INSERT INTO prescriptions (patient_id, doctor_id, medical_record_id, medication_details)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
`
	return pr.database.QueryRowContext(c, query,
		prescription.PatientID,
		prescription.DoctorID,
		prescription.MedicalRecordID,
		prescription.MedicationDetails,
	).Scan(&prescription.ID, &prescription.CreatedAt)
}

func (pr *prescriptionRepository) Fetch(c context.Context) ([]domain.Prescription, error) {
	query := `
	SELECT id, patient_id, doctor_id, medical_record_id, medication_details, created_at
    FROM prescriptions
`
	rows, err := pr.database.QueryContext(c, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching prescription: %w", err)
	}
	defer rows.Close()

	var prescriptions []domain.Prescription
	for rows.Next() {
		var prescription domain.Prescription
		if err := rows.Scan(
			&prescription.ID,
			&prescription.PatientID,
			&prescription.DoctorID,
			&prescription.MedicalRecordID,
			&prescription.MedicationDetails,
			&prescription.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning prescription: %w", err)
		}

		prescriptions = append(prescriptions, prescription)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating prescription: %w", err)
	}

	return prescriptions, nil
}

func (pr *prescriptionRepository) FetchByID(c context.Context, id uuid.UUID) (*domain.Prescription, error) {
	query := `
	SELECT id, patient_id, doctor_id, medical_record_id, medication_details, created_at
	FROM prescriptions
	WHERE id = $1
`
	prescription := &domain.Prescription{}
	err := pr.database.QueryRowContext(c, query, id).Scan(
		&prescription.ID,
		&prescription.PatientID,
		&prescription.DoctorID,
		&prescription.MedicalRecordID,
		&prescription.MedicationDetails,
		&prescription.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error fetching prescription: %w", err)
	}

	return prescription, nil
}

func (pr *prescriptionRepository) Update(c context.Context, prescription *domain.Prescription) error {
	query := `
	UPDATE prescriptions
	SET patient_id = $1, doctor_id = $2, medical_record_id = $3, medication_details = $4
	WHERE id = $5
`
	result, err := pr.database.ExecContext(c, query,
		&prescription.PatientID,
		&prescription.DoctorID,
		&prescription.MedicalRecordID,
		&prescription.MedicationDetails,
		&prescription.ID,
	)

	if err != nil {
		return fmt.Errorf("Error updating prescription: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("prescription with ID %s not found", prescription.ID)
	}

	return nil
}

func (pr *prescriptionRepository) Delete(c context.Context, id uuid.UUID) error {
	query := `
	DELETE FROM prescriptions
	WHERE id = $1
	`
	result, err := pr.database.ExecContext(c, query, id)
	if err != nil {
		return fmt.Errorf("error deleting prescription: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("prescription with ID %s not found", id)
	}

	return nil

}
