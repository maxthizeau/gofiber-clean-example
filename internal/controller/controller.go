package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/service"
)

type UserController struct {
	service.UserService
	// configuration.Config
}

type QuestionController struct {
	service.QuestionService
	// configuration.Config
}

type Controllers struct {
	UserController
	QuestionController
}

func NewControllers(services *service.Services) *Controllers {
	return &Controllers{
		UserController:     *NewUserController(&services.UserService),
		QuestionController: *NewQuestionController(&services.QuestionService),
	}
}

func (c *Controllers) ServeRoutes(app *fiber.App) {
	c.QuestionController.Route(app)
	c.UserController.Route(app)
}
