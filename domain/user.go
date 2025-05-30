package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	AdminRole UserRole = "admin"
	DoctorRole UserRole = "doctor"
	PatientRole UserRole = "patient"
)

type User struct {
	ID        uuid.UUID `json:"user_id"`
	Username  string    `json:"username"`
	Email 	  string    `json:"email"`
	Password  string    `json:"password"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id uuid.UUID) (User, error)
}