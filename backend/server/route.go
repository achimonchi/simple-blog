package server

import (
	"log"
	"net/http"
	"simple-blog/config"
	"simple-blog/server/controller"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	app        *config.ConfigApp
	router     *httprouter.Router
	controller *controller.Controller
}

func NewRoute(app *config.ConfigApp, router *httprouter.Router, controller *controller.Controller) *Route {
	return &Route{app, router, controller}
}

func (r *Route) StartServer() {
	log.Println("server running at port", r.app.Port)

	r.router.POST("/users/registration", r.controller.User.Registration)
	r.router.POST("/users/login", r.controller.User.Login)

	http.ListenAndServe(r.app.Port, r.router)
}
