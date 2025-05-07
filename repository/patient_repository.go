package repository

import (
	"context"
	"database/sql"
	"fmt"
	"hms-api/domain"

	"github.com/google/uuid"
)

type patientRepository struct {
	database *sql.DB
}

func NewPatientRepository(db *sql.DB) domain.PatientRepository {
		return &patientRepository{
		database: db,
		}
}

func (pr *patientRepository) Create(c context.Context, patient *domain.Patient) error {
	query := `
		INSERT INTO patients (user_id, cpf, date_birth, phone, address)
		VALUES ($1, $2, $3, $4, $5) RETURNING id
	`
	err := pr.database.QueryRowContext(c, query, patient.UserId, patient.CPF, patient.DateBirth, patient.Phone, patient.Address).Scan(&patient.ID)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return err
	}
	
	return nil
}

func (pr *patientRepository) Fetch(c context.Context) ([]domain.Patient, error) {
	query := `
		SELECT id, user_id, cpf, date_birth, phone, address, created_at
		FROM patients
	`
	rows, err := pr.database.QueryContext(c, query)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return []domain.Patient{}, err
	}
	defer rows.Close()

	var patients []domain.Patient

	for rows.Next(){
		var patient domain.Patient
		err = rows.Scan(
			&patient.ID,
			&patient.UserId,
			&patient.CPF,
			&patient.DateBirth,
			&patient.Phone,
			&patient.Address,
			&patient.CreatedAt,
		)
			if err != nil {
		fmt.Println("Error scanning row:", err)
		return []domain.Patient{}, err
	}

	patients = append(patients, patient)
	}

	if err = rows.Err(); err != nil{
		return nil, err
	}

	return patients, nil
}

func (pr *patientRepository) FetchByID(c context.Context, id uuid.UUID) (domain.Patient, error) {
	var patient domain.Patient
	query := `
		SELECT id, user_id, cpf, date_birth, phone, address, created_at
		FROM patients WHERE id = $1
	`

	err := pr.database.QueryRowContext(c, query, id).Scan(
		&patient.ID,
		&patient.UserId,
		&patient.CPF,
		&patient.DateBirth,
		&patient.Phone,
		&patient.Address,
		&patient.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Patient{}, nil
		}
		return domain.Patient{}, err
	}

	return patient, nil
}

func (pr *patientRepository) Update(c context.Context, patient *domain.Patient) error {
	query := `
		UPDATE patients
		SET cpf = $1, date_birth = $2, phone = $3, address = $4
		WHERE id = $5
		`
		_, err := pr.database.ExecContext(c, query, patient.CPF, patient.DateBirth, patient.Phone, patient.Address, patient.ID)

		if err != nil {
		fmt.Println("Error executing update:", err)
		return err
		}

		return nil
}

func (pr *patientRepository) Delete(c context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM patients
		WHERE id = $1
	`
	_, err := pr.database.ExecContext(c, query, id)
	if err != nil {
		fmt.Println("Error executing delete:", err)
		return err
	}
	
	return nil
}