package repository

import (
	"database/sql"
	"fmt"
	"hms-api/model"

	"github.com/google/uuid"
)

type AppointmentRepository struct {
	connection *sql.DB	
}

func NewAppointmentRepository(db *sql.DB) AppointmentRepository {
	return AppointmentRepository{
		connection: db,
	}
}

func (pr *AppointmentRepository) GetAppointments() ([]model.Appointment, error){
	query := "SELECT id, appointment_date, status, notes, created_at, updated_at FROM appointments"
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.Appointment{}, err
	}

	var appointmentList []model.Appointment
	var appointmentObj model.Appointment

	for rows.Next(){
		err = rows.Scan(
			&appointmentObj.ID,
			&appointmentObj.AppointmentDate,
			&appointmentObj.Status,
			&appointmentObj.Notes,
			&appointmentObj.CreatedAt,
			&appointmentObj.UpdatedAt,
		)

		if(err != nil){
			fmt.Println("Error executing query:", err)
			return []model.Appointment{}, err
		}

		appointmentList = append(appointmentList, appointmentObj)
	}

	rows.Close()

	return appointmentList, nil
}

func (pr *AppointmentRepository) GetAppointmentById(appintment_id uuid.UUID) (*model.Appointment, error){
	query, err := pr.connection.Prepare("SELECT id, appointment_date, status, notes, created_at, updated_at FROM appointments WHERE id = $1")
	if err != nil {
		fmt.Println("Error preparing query:", err)
		return nil, err
	}

	var appointment model.Appointment
	err = query.QueryRow(appintment_id).Scan(
		&appointment.ID,
		&appointment.AppointmentDate,
		&appointment.Status,
		&appointment.Notes,
		&appointment.CreatedAt,
		&appointment.UpdatedAt,
	)

	if(err != nil){
		if(err == sql.ErrNoRows){
			fmt.Println("No appointment found with the given ID")
			return nil, nil
		}
		fmt.Println("Error executing query:", err)
		return nil, err
	}

	query.Close()
	return &appointment, nil
}

func (pr *AppointmentRepository) CreateAppointment(appointment model.Appointment) (uuid.UUID, error) {
	var id uuid.UUID
	query, err := pr.connection.Prepare("INSERT INTO appointments" +
	"(appointment_date, status, notes)" +
	"VALUES ($1, $2, $3) RETURNING id")

	if err != nil {
		fmt.Println("Error preparing query:", err)
		return uuid.UUID{}, err
	}

	err = query.QueryRow(appointment.AppointmentDate, appointment.Status, appointment.Notes).Scan(&id)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return uuid.UUID{}, err
	}

	query.Close()
	return id, nil

}

func (pr *AppointmentRepository) UpdateAppointment(appointment_id uuid.UUID, appointment model.Appointment) (uuid.UUID, error){
	query, err := pr.connection.Prepare(`
		DELETE FROM appointments
		WHERE id = $1
		RETURNING id, appointment_date, status, notes, created_at, updated_at
	`)

	if err != nil {
		fmt.Println("Error preparing query:", err)
		return uuid.UUID{}, err
	}

	err = query.QueryRow(appointment.AppointmentDate, appointment.Status, appointment.Notes, appointment_id).Scan(&appointment_id)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return uuid.UUID{}, err
	}

	query.Close()
	return appointment_id, nil
}

func (ar *AppointmentRepository) DeleteAppointment(appointment_id uuid.UUID) (*model.Appointment, error) {
	query, err := ar.connection.Prepare("DELETE FROM appointments WHERE id = $1")

	if err != nil {
		fmt.Println("Error preparing query:", err)
		return &model.Appointment{}, err
	}

	var appointment = model.Appointment{}
	err = query.QueryRow(appointment_id).Scan(&appointment.ID, &appointment.AppointmentDate, &appointment.Status, &appointment.Notes, &appointment.CreatedAt, &appointment.UpdatedAt)

	if(err != nil){
		if(err == sql.ErrNoRows){
			fmt.Println("No appointment found with the given ID")
			return nil, nil
		}
		fmt.Println("Error executing query:", err)
		return nil, err
	}

	query.Close()
	return &appointment, nil

}