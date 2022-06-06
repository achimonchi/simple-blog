package service

import (
	"database/sql"
	"net/http"
	"simple-blog/data/params"
	"simple-blog/helper"
	"simple-blog/server/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) RegisterUser(req *params.UserCreate) *params.Response {
	user := req.ParseToModel()

	userDB, err := u.repo.GetUserByEmail(user.Email)
	if err != nil {
		if err != sql.ErrNoRows {
			return &params.Response{
				Success: false,
				Status:  http.StatusInternalServerError,
			}
		}
	}

	if userDB != nil {
		return &params.Response{
			Success: false,
			Status:  http.StatusConflict,
		}
	}

	err = u.repo.CreateUser(user)
	if err != nil {
		return &params.Response{
			Success: false,
			Status:  http.StatusInternalServerError,
		}

	}
	return &params.Response{
		Success: true,
		Status:  http.StatusCreated,
	}
}

func (u *UserService) LoginUser(req *params.UserLogin) *params.Response {
	user, err := u.repo.GetUserByEmail(req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return &params.Response{
				Success: false,
				Status:  http.StatusNotFound,
			}
		} else {
			return &params.Response{
				Success: false,
				Status:  http.StatusInternalServerError,
			}
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return &params.Response{
			Success: false,
			Status:  http.StatusUnauthorized,
		}
	}

	token := helper.GenerateToken(uint(user.Id), user.Email)

	return &params.Response{
		Success: true,
		Status:  http.StatusOK,
		Token:   token,
	}
}
