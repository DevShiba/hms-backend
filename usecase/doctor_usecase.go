package usecase

import (
	"hms-api/model"
	"hms-api/repository"

	"github.com/google/uuid"
)

type DoctorUsecase struct {
	repository repository.DoctorRepository
}

func NewDoctorUsecase(repo repository.DoctorRepository) DoctorUsecase {
	return DoctorUsecase{
		repository: repo,
	}
}

func (du *DoctorUsecase) GetDoctors() ([]model.Doctor, error) {
	doctorList, err := du.repository.GetDoctors()
	if err != nil {
		return []model.Doctor{}, err
	}

	return doctorList, nil
}

func (du *DoctorUsecase) GetDoctorById(doctor_id uuid.UUID) (*model.Doctor, error){
	doctor, err := du.repository.GetDoctorById(doctor_id)
	if err != nil {
		return nil, err
	}

	return doctor, nil
}

func (du *DoctorUsecase) CreateDoctor(doctor model.Doctor) (*model.Doctor, error){
    createdDoctor, err := du.repository.CreateDoctor(doctor)
    if err != nil {
        return nil, err
    }
    
    return createdDoctor, nil
}

func (du *DoctorUsecase) UpdateDoctor(doctor_id uuid.UUID, doctor model.Doctor) (*model.Doctor, error){
	updatedDoctor, err := du.repository.UpdateDoctor(doctor_id, doctor)
	if err != nil {
		return nil, err
	}

	return updatedDoctor, nil
}

func (du *DoctorUsecase) DeleteDoctor(doctor_id uuid.UUID) error {
    err := du.repository.DeleteDoctor(doctor_id)
    if err != nil {
        return err
    }

    return nil
}