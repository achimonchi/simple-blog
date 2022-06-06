package controller

import (
	"encoding/json"
	"net/http"
	"simple-blog/data/params"
)

type Controller struct {
	User *UserController
}

func NewController(User *UserController) *Controller {
	return &Controller{User}
}

type base struct {
	rw       http.ResponseWriter
	response *params.Response
}

func writeResponse(rw http.ResponseWriter, response *params.Response) *base {

	rw.WriteHeader(response.Status)

	var b = base{
		rw:       rw,
		response: response,
	}

	return &b
}

func (b *base) build() {
	b.rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(b.rw).Encode(b.response)
}
