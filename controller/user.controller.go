package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxthizeau/gofiber-clean-boilerplate/configuration"
	"github.com/maxthizeau/gofiber-clean-boilerplate/exception"
	"github.com/maxthizeau/gofiber-clean-boilerplate/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/service"
)

type UserController struct {
	service.UserService
	configuration.Config
}

func NewUserController(userService *service.UserService, config configuration.Config) *UserController {
	return &UserController{UserService: *userService, Config: config}
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/v1/api/sign-up", controller.SignUp)
	app.Get("/v1/api/user", controller.FindAll)
}

func (controller UserController) SignUp(c *fiber.Ctx) error {
	var request model.UserAuthenticationModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	result := controller.UserService.SignUp(c.Context(), request)

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

func (controller UserController) FindAll(c *fiber.Ctx) error {
	result := controller.UserService.FindAll(c.Context())
	return c.Status(200).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
