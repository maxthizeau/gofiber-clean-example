package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/service"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
)

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: *userService}
}

func (controller UserController) generateTokenForUser(user entity.User) string {
	roles := []string{}
	for _, r := range user.UserRoles {
		roles = append(roles, r.Role)
	}
	return auth.GenerateJwtToken(user.Id, roles, "SECRET", 60)
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/v1/api/sign-up", controller.SignUp)
	app.Post("/v1/api/auth", controller.Authenticate)
	app.Get("/v1/api/user", controller.FindAll)
}

func (controller UserController) SignUp(c *fiber.Ctx) error {
	var request model.UserSignupModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	user := controller.UserService.SignUp(c.Context(), request)

	result := model.AuthModel{
		User: model.UserModel{
			Username: user.Username,
			Email:    user.Email,
		},
		Token: controller.generateTokenForUser(user),
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

func (controller UserController) Authenticate(c *fiber.Ctx) error {
	var request model.UserLoginModel

	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	user := controller.UserService.Authenticate(c.Context(), request)

	result := model.AuthModel{
		User: model.UserModel{
			Username: user.Username,
			Email:    user.Email,
		},
		Token: controller.generateTokenForUser(user),
	}

	return c.Status(200).JSON(model.GeneralResponse{
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
