package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/maxthizeau/gofiber-clean-boilerplate/configuration"
	"github.com/maxthizeau/gofiber-clean-boilerplate/controller"
	"github.com/maxthizeau/gofiber-clean-boilerplate/exception"
	repository "github.com/maxthizeau/gofiber-clean-boilerplate/repository/impl"
	service "github.com/maxthizeau/gofiber-clean-boilerplate/service/impl"
)

func main() {
	log.Println("Work in progres...")
	config := configuration.New()
	database := configuration.NewDatabase(config)

	// repository
	userRepository := repository.NewUserRepositoryImpl(database)

	// service
	userService := service.NewUserServiceImpl(&userRepository)

	// controller
	userController := controller.NewUserController(&userService, config)

	// fiber
	app := fiber.New(configuration.NewFiberConfiguration())

	// route
	userController.Route(app)

	err := app.Listen(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)
}
