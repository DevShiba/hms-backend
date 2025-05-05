package repository

import (
	"context"
	"database/sql"
	"fmt"
	"hms-api/domain"

	"github.com/google/uuid"
)

type doctorRepository struct {
	database *sql.DB
}

func NewDoctorRepository(db *sql.DB) domain.DoctorRepository {
	return &doctorRepository{
		database: db,
	}
}

func (dr *doctorRepository) Create(c context.Context, doctor *domain.Doctor) error {
	query := "INSERT INTO doctors (user_id, crm, specialty) VALUES ($1, $2, $3) RETURNING id"
	err := dr.database.QueryRowContext(c, query, doctor.UserId, doctor.CRM, doctor.Specialty).Scan(&doctor.ID)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return err
	}
	return nil
}

func (dr *doctorRepository) Fetch(c context.Context) ([]domain.Doctor, error) {
	query := "SELECT id, user_id, crm, specialty, created_at FROM doctors"
	rows, err := dr.database.QueryContext(c, query)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return []domain.Doctor{}, err
	}
	defer rows.Close()

	var doctors []domain.Doctor

	for rows.Next() {
		var doctor domain.Doctor
		err = rows.Scan(
			&doctor.ID,
			&doctor.UserId,
			&doctor.CRM,
			&doctor.Specialty,
			&doctor.CreatedAt,
		)

		if err != nil {
			fmt.Println("Error scanning row:", err)
			return []domain.Doctor{}, err
		}

		doctors = append(doctors, doctor)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return doctors, nil
}

func (dr *doctorRepository) FetchByID(c context.Context, id uuid.UUID) (domain.Doctor, error) {
	var doctor domain.Doctor
	query := "SELECT id, user_id, crm, specialty, created_at FROM doctors WHERE id = $1"
	
	err := dr.database.QueryRowContext(c, query, id).Scan(
		&doctor.ID,
		&doctor.UserId,
		&doctor.CRM,
		&doctor.Specialty,
		&doctor.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Doctor{}, nil
		}
		return domain.Doctor{}, err
	}

	return doctor, nil
}

func (dr *doctorRepository) Update(c context.Context, doctor *domain.Doctor) error {
	query := "UPDATE doctors SET crm = $1, specialty = $2 WHERE id = $3"
	_, err := dr.database.ExecContext(c, query, doctor.CRM, doctor.Specialty, doctor.ID)
	
	if err != nil {
		fmt.Println("Error executing update:", err)
		return err
	}

	return nil
}

func (dr *doctorRepository) Delete(c context.Context, id uuid.UUID) error {
	query := "DELETE FROM doctors WHERE id = $1"
	_, err := dr.database.ExecContext(c, query, id)
	
	if err != nil {
		fmt.Println("Error executing delete:", err)
		return err
	}

	return nil
}