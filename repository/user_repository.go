package repository

import (
	"database/sql"
	"fmt"
	"hms-api/model"

	"github.com/google/uuid"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository {
		connection: db,
	}
}

func (ur *UserRepository) RegisterUser(user model.User) (uuid.UUID, error)  {
	var id uuid.UUID
	query, err := ur.connection.Prepare("INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id")
	
	if err != nil {
		fmt.Println("Error preparing query:", err)
		return uuid.UUID{}, err
	}

	err = query.QueryRow(user.Username, user.Email, user.Password, user.Role).Scan(&user.ID)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return uuid.UUID{}, err
	}

	query.Close()
	return id, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (model.User, error){
	var u model.User

	query := `
		SELECT id, username, email, password, role, created_at, updated_at
		FROM users WHERE email = $1
	`

	err := ur.connection.QueryRow(query, email).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.Password,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, nil
		}
		return model.User{}, err
	}

	return u, nil
}

