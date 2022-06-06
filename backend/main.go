package main

import (
	"fmt"
	"os"
	"simple-blog/config"
	"simple-blog/constants"
	"simple-blog/database"
	"simple-blog/server"
	"simple-blog/server/controller"
	"simple-blog/server/repository"
	"simple-blog/server/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("Server running ...")
	conf := initConfig()

	db := database.GetConnection(conf.Database)
	if db != nil {
		fmt.Println("db connected")
	}

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	controller := controller.NewController(userController)

	router := httprouter.New()
	route := server.NewRoute(conf.App, router, controller)

	route.StartServer()
}

func initConfig() *config.Config {
	var config config.Config
	config.Database = initDB()
	config.App = initApp()
	return &config
}

func initDB() *config.ConfigDB {
	dbHost := os.Getenv(constants.POSTGRES_HOST)
	dbPort := os.Getenv(constants.POSTGRES_PORT)
	dbUser := os.Getenv(constants.POSTGRES_USER)
	dbPass := os.Getenv(constants.POSTGRES_PASS)
	dbName := os.Getenv(constants.POSTGRES_DBNAME)

	maxIdleConns, _ := strconv.Atoi(os.Getenv(constants.POSTGRES_MAX_IDLE_CONNS))
	maxOpenConns, _ := strconv.Atoi(os.Getenv(constants.POSTGRES_MAX_OPEN_CONNS))
	maxIdleTimeConns, _ := strconv.Atoi(os.Getenv(constants.POSTGRES_MAX_IDLE_TIME_CONNS))
	maxLifeTimeConns, _ := strconv.Atoi(os.Getenv(constants.POSTGRES_MAX_LIFE_TIME_CONNS))

	var dbConfig = config.ConfigDB{
		Host:         dbHost,
		Port:         dbPort,
		User:         dbUser,
		Pass:         dbPass,
		DBName:       dbName,
		SSLMode:      "disable",
		MaxIdleConns: maxIdleConns,
		MaxOpenConns: maxOpenConns,
		MaxLifeTime:  maxLifeTimeConns,
		MaxIdleTime:  maxIdleTimeConns,
	}

	return &dbConfig
}

func initApp() *config.ConfigApp {
	port := os.Getenv(constants.APP_PORT)

	return &config.ConfigApp{
		Port: port,
	}
}
