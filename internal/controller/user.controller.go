package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/service"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth/role"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
)

func NewUserController(userService *service.UserService, authManager auth.AuthManager) *UserController {
	return &UserController{UserService: *userService, Auth: authManager}

}

func (controller UserController) generateTokenForUser(user entity.User) string {
	roles := role.RolesList{}
	for _, r := range user.UserRoles {
		roles = append(roles, role.RoleEnum(r.Role))
	}
	return controller.Auth.GenerateJwtToken(user.Id, roles)
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/v1/api/sign-up", controller.SignUp)
	app.Post("/v1/api/auth", controller.Authenticate)
	app.Get("/v1/api/user", controller.FindAll)
}

func (controller UserController) SignUp(c *fiber.Ctx) error {
	var request model.UserSignupInput
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	user := controller.UserService.SignUp(c.Context(), request)

	result := model.Auth{
		User:  model.NewUserFromEntity(user),
		Token: controller.generateTokenForUser(user),
	}

	response := model.NewSuccessResponse(result)

	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller UserController) Authenticate(c *fiber.Ctx) error {
	var request model.UserLoginInput

	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	user := controller.UserService.Authenticate(c.Context(), request)

	result := model.Auth{
		User:  model.NewUserFromEntity(user),
		Token: controller.generateTokenForUser(user),
	}

	response := model.NewSuccessResponse(result)
	return c.Status(200).JSON(response)
}

func (controller UserController) FindAll(c *fiber.Ctx) error {
	result := controller.UserService.FindAll(c.Context())
	response := model.NewSuccessResponse(result)
	return c.Status(200).JSON(response)
}
