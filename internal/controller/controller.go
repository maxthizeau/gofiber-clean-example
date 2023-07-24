package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/middleware"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/service"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth"
)

type UserController struct {
	service.UserService
	Middlewares middleware.MiddlewareManager
	Auth        auth.AuthManager
}

type QuestionController struct {
	service.QuestionService
	Middlewares middleware.MiddlewareManager
	Auth        auth.AuthManager
}

type Controllers struct {
	UserController
	QuestionController
}

func NewControllers(services *service.Services, middlewareManager middleware.MiddlewareManager, authManager auth.AuthManager) *Controllers {
	return &Controllers{
		UserController:     *NewUserController(&services.UserService, authManager),
		QuestionController: *NewQuestionController(&services.QuestionService, middlewareManager, authManager),
	}
}

func (c *Controllers) ServeRoutes(app *fiber.App) {
	c.QuestionController.Route(app)
	c.UserController.Route(app)
}
