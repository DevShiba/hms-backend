package repository

import (
	"context"
	"database/sql"
	"fmt"
	"hms-api/domain"

	"github.com/google/uuid"
)

type appointmentRepository struct {
	database *sql.DB
}

func NewAppointmentRepository(db *sql.DB) domain.AppointmentRepository {
	return &appointmentRepository{
		database: db,
	}
}

func (ar *appointmentRepository) Create(c context.Context, appointment *domain.Appointment) error {
	query := `
		INSERT INTO appointments (patient_id, doctor_id, appointment_date, status, notes)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := ar.database.QueryRowContext(c, query, appointment.PatientID, appointment.DoctorID, appointment.AppointmentDate, appointment.Status, appointment.Notes).Scan(&appointment.ID)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return err
	}

	return nil
}

func (ar *appointmentRepository) Fetch(c context.Context) ([]domain.Appointment, error) {
	query := `
		SELECT id, patient_id, doctor_id, appointment_date, status, notes, created_at, updated_at
		FROM appointments
	`
	rows, err := ar.database.QueryContext(c, query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []domain.Appointment{}, err
	}

	defer rows.Close()

	var appointments []domain.Appointment

	for rows.Next() {
		var appointment domain.Appointment
		err = rows.Scan(
			&appointment.ID,
			&appointment.PatientID,
			&appointment.DoctorID,
			&appointment.AppointmentDate,
			&appointment.Status,
			&appointment.Notes,
			&appointment.CreatedAt,
			&appointment.UpdatedAt,
		)

		if err != nil {
			fmt.Println("Error scanning row:", err)
			return []domain.Appointment{}, err
		}

		appointments = append(appointments, appointment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return appointments, nil
}

func (ar *appointmentRepository) FetchByID(c context.Context, id uuid.UUID) (domain.Appointment, error) {
	var appointment domain.Appointment
	query := `
		SELECT id, patient_id, doctor_id, appointment_date, status, notes, created_at, updated_at
		FROM appointments
		WHERE id = $1
	`

	err := ar.database.QueryRowContext(c, query, id).Scan(
		&appointment.ID,
		&appointment.PatientID,
		&appointment.DoctorID,
		&appointment.AppointmentDate,
		&appointment.Status,
		&appointment.Notes,
		&appointment.CreatedAt,
		&appointment.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Appointment{}, nil
		}
		return domain.Appointment{}, err
	}

	return appointment, nil
}

func (ar *appointmentRepository) Update(c context.Context, appointment *domain.Appointment) error {
	query := `
		UPDATE appointments
		SET patient_id = $1, doctor_id = $2, appointment_date = $3, status = $4, notes = $5
		WHERE id = $6
	`

	_, err := ar.database.ExecContext(c, query, appointment.PatientID, appointment.DoctorID, appointment.AppointmentDate, appointment.Status, appointment.Notes, appointment.ID)

	if err != nil {
		fmt.Println("Error executing update:", err)
		return err
	}

	return nil
}

func (ar *appointmentRepository) Delete(c context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM appointments
		WHERE id = $1
	`
	_, err := ar.database.ExecContext(c, query, id)

	if err != nil {
		fmt.Println("Error executing delete:", err)
		return err
	}

	return nil
}
