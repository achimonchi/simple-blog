package params

import (
	"simple-blog/data/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserCreate struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *UserCreate) ParseToModel() *models.User {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	return &models.User{
		FullName:  u.FullName,
		Password:  string(hashed),
		Email:     u.Email,
		CreatedAt: time.Now(),
	}
}

type UserLogin struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
