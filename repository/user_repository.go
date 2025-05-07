package repository

import (
	"context"
	"database/sql"
	"fmt"
	"hms-api/domain"

	"github.com/google/uuid"
)

type userRepository struct {
	database *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
		database:  db,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
    query := `
			INSERT INTO users (username, email, password, role) 
			VALUES ($1, $2, $3, $4) RETURNING id
		`

    err := ur.database.QueryRowContext(c, query, user.Username, user.Email, user.Password, user.Role).Scan(&user.ID)
    if err != nil {
        fmt.Println("Error executing insert:", err)
        return err
    }
    return nil
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	query := `
		SELECT id, username, email, password, role, created_at, updated_at 
		FROM users
	`

    rows, err := ur.database.QueryContext(c, query)
    if err != nil {
        fmt.Println("Error executing query:", err)
        return nil, err
    }
    defer rows.Close()

    var users []domain.User

    for rows.Next() {
        var user domain.User
        err := rows.Scan(
            &user.ID,
            &user.Username,
            &user.Email,
            &user.Password,
            &user.Role,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
        if err != nil {
            fmt.Println("Error scanning row:", err)
            return nil, err
        }
        users = append(users, user)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
    var user domain.User
    query := `
        SELECT id, username, email, password, role, created_at, updated_at
        FROM users WHERE email = $1
    `
    err := ur.database.QueryRowContext(c, query, email).Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &user.Password,
        &user.Role,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            return domain.User{}, fmt.Errorf("user not found with email: %s", email)
        }
        return domain.User{}, err
    }
    return user, nil
}

func (ur *userRepository) GetByID(c context.Context, id uuid.UUID) (domain.User, error){
	query := `
		SELECT id, username, email, password, role, created_at, updated_at
		FROM users WHERE id = $1
	`

	var user domain.User
	err := ur.database.QueryRowContext(c, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, fmt.Errorf("user not found with ID: %s", id)
		}
		return domain.User{}, err
	}

	return user, nil
}


