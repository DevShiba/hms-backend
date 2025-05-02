package repository

import (
	"database/sql"
	"fmt"
	"hms-api/model"

	"github.com/google/uuid"
)

type DoctorRepository struct {
	connection *sql.DB
}

func NewDoctorRepository(db *sql.DB) DoctorRepository {
	return DoctorRepository{
		connection: db,
	}
}

func (dr *DoctorRepository) GetDoctors() ([]model.Doctor, error){
	query := "SELECT id, user_id, crm, specialty, created_at FROM doctors"
	rows, err := dr.connection.Query(query)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.Doctor{}, err
	}

	var doctorList []model.Doctor
	var doctorObj model.Doctor

	for rows.Next(){
		err = rows.Scan(
			&doctorObj.ID,
			&doctorObj.UserId,
			&doctorObj.CRM,
			&doctorObj.Specialty,
			&doctorObj.CreatedAt,
		)

		if(err != nil){
			fmt.Println("Error executing query:", err)
			return []model.Doctor{}, err
		}

		doctorList = append(doctorList, doctorObj)
	}

	rows.Close()

	return doctorList, nil
}

func (dr *DoctorRepository) CreateDoctor(doctor model.Doctor) (*model.Doctor, error) {
	query := "INSERT INTO doctors (user_id, crm, specialty) VALUES ($1, $2, $3) RETURNING id"
	rows, err := dr.connection.Query(query, doctor.UserId, doctor.CRM, doctor.Specialty)
	if err != nil {
		fmt.Println("Error preparing query:", err)
		return nil, err
	}

	var doctorObj model.Doctor
	for rows.Next() {
		err = rows.Scan(
			&doctorObj.ID,
			&doctorObj.UserId,
			&doctorObj.CRM,
			&doctorObj.Specialty,
			&doctorObj.CreatedAt,
		)
	}

	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}

	rows.Close()

	return &doctorObj, nil

}

func (dr *DoctorRepository) GetDoctorById(doctor_id uuid.UUID) (*model.Doctor, error){
	query := "SELECT id, user_id, crm, specialty, created_at FROM doctors WHERE id = $1"	
	rows, err := dr.connection.Query(query, doctor_id)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}

	var doctorObj model.Doctor

	for rows.Next(){
		err = rows.Scan(
			&doctorObj.ID,
			&doctorObj.UserId,
			&doctorObj.CRM,
			&doctorObj.Specialty,
			&doctorObj.CreatedAt,
		)

		if(err != nil){
			fmt.Println("Error executing query:", err)
			return nil, err
		}
	}

		rows.Close()

	return &doctorObj, nil
}

func (dr *DoctorRepository) UpdateDoctor(doctor_id uuid.UUID, doctor model.Doctor) (*model.Doctor, error) {
	query := "UPDATE doctors SET crm = $1, specialty = $2 WHERE id = $3"
	rows, err := dr.connection.Query(query, doctor.CRM, doctor.Specialty, doctor.ID)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}

	var doctorObj model.Doctor


	for rows.Next(){
		err = rows.Scan(
			&doctorObj.ID,
			&doctorObj.UserId,
			&doctorObj.CRM,
			&doctorObj.Specialty,
			&doctorObj.CreatedAt,
		)

		if(err != nil){
			fmt.Println("Error executing query:", err)
			return nil, err
		}

	}
	
	rows.Close()

	return &doctorObj, nil
}

func (dr *DoctorRepository) DeleteDoctor(doctor_id uuid.UUID) error {
	query := "DELETE FROM doctors WHERE id = $1"
	_, err := dr.connection.Query(query, doctor_id)
	
	if err != nil {
		fmt.Println("Error executing query:", err)
		return  err
	}

	return nil
}