package controller

import (
	"encoding/json"
	"net/http"
	"simple-blog/data/params"
	"simple-blog/server/service"

	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) Registration(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req params.UserCreate

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(rw, &params.Response{
			Success: false,
			Status:  http.StatusBadRequest,
		}).build()
		return
	}

	resp := u.userService.RegisterUser(&req)

	writeResponse(rw, resp).build()
}

func (u *UserController) Login(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req params.UserLogin
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(rw, &params.Response{
			Success: false,
			Status:  http.StatusBadRequest,
		}).build()
		return
	}

	resp := u.userService.LoginUser(&req)
	writeResponse(rw, resp).build()
}
