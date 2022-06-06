package repository

import (
	"database/sql"
	"simple-blog/data/models"
)

type UserRepo interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users(full_name, email, password, created_at)
		VALUES($1,$2,$3,$4)
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.FullName, user.Email, user.Password, user.CreatedAt)

	return err
}

func (r *userRepo) GetUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT 
			id, full_name, email, password, created_at
		FROM 
			users
		WHERE email=$1
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(email)

	var user models.User
	err = row.Scan(
		&user.Id, &user.FullName, &user.Email, &user.Password, &user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
