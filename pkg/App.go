package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"log"
	"user-service/pkg/db"
	"user-service/pkg/handler"
	"user-service/pkg/middleware"
	"user-service/pkg/service"
)

func main() {
	initConfig()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))
	e.Use(middleware_user_service.LogRequest)
	err, repFactory := db.InitRepositoryFactory(new(db.ConnectionFactory))
	if err != nil {
		log.Fatal(err)
	}
	groupService := new(service.GroupService)
	userService := new(service.UserService)
	err = service.InitializeAllServices(repFactory, groupService, userService)
	if err != nil {
		log.Fatal(err)
	}

	groupHandler := &handler.GroupHandler{E: e, Service: groupService}
	userHandler := &handler.UserHandler{E: e, Service: userService}
	err = handler.InitializeAllHandlers(e, groupHandler, userHandler)
	if err != nil {
		log.Fatal(err)
	}

	e.Logger.Fatal(e.Start(viper.Get("server.host").(string) + viper.Get("server.port").(string)))
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
