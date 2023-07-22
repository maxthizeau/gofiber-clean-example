package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

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
	questionRepository := repository.NewQuestionRepositoryImpl(database)
	// answerRepository := repository.NewAnswerRepositoryImpl(database)

	// service
	userService := service.NewUserServiceImpl(&userRepository, config)
	questionService := service.NewQuestionServiceImpl(&questionRepository)

	// controller
	userController := controller.NewUserController(&userService, config)
	questionController := controller.NewQuestionController(&questionService, config)

	// fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())

	// route
	userController.Route(app)
	questionController.Route(app)

	err := app.Listen(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)
}
