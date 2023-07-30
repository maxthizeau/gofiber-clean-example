package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/middleware"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/service"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth/role"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
)

func NewUserController(userService *service.UserService, middleware middleware.MiddlewareManager, authManager auth.AuthManager) *UserController {
	return &UserController{UserService: *userService, Middlewares: middleware, Auth: authManager}

}

func (controller UserController) generateTokenForUser(user entity.User) string {
	roles := role.RolesList{}
	for _, r := range user.UserRoles {
		roles = append(roles, role.RoleEnum(r.Role))
	}
	return controller.Auth.GenerateJwtToken(user.Id, roles)
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/v1/api/signup", controller.SignUp)
	app.Post("/v1/api/login", controller.Authenticate)
	app.Post("/v1/api/refresh", controller.Middlewares.AuthenticateJWT(), controller.Refresh)
}

// @Summary User Sign up
// @Tags user
// @Description Create user and return JWT token + user associated
// @ModuleID SignUp
// @Accept  json
// @Produce  json
// @Param input body model.UserSignupInput true "sign up info"
// @Success 200 {object} model.GeneralResponse{data=model.Auth}
// @Failure 400 {object} model.GeneralResponse{data=[]common.ValidationResponse}
// @Router /signup [post]
func (controller UserController) SignUp(c *fiber.Ctx) error {
	var request model.UserSignupInput
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	user := controller.UserService.SignUp(c.Context(), request)

	userModel := model.NewUserFromEntity(user)
	result := model.Auth{
		User:  &userModel,
		Token: controller.generateTokenForUser(user),
	}

	response := model.NewSuccessResponse(result)

	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary User Login
// @Tags user
// @Description Verify provided credentials and return JWT token + user associated
// @ModuleID Authenticate
// @Accept  json
// @Produce  json
// @Param input body model.UserLoginInput true "sign in info"
// @Success 200 {object} model.GeneralResponse{data=model.Auth}
// @Failure 400 {object} model.GeneralResponse{data=[]common.ValidationResponse}
// @Router /login [post]
func (controller UserController) Authenticate(c *fiber.Ctx) error {
	var request model.UserLoginInput

	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	user := controller.UserService.Authenticate(c.Context(), request)
	userModel := model.NewUserFromEntity(user)

	result := model.Auth{
		User:  &userModel,
		Token: controller.generateTokenForUser(user),
	}

	response := model.NewSuccessResponse(result)
	return c.Status(200).JSON(response)
}

// @Summary User Refresh
// @Tags user
// @Description Verify the Authorization token, refresh it if valid, return null user otherwise
// @ModuleID Refresh
// @Accept  json
// @Produce  json
// @Success 200 {object} model.GeneralResponse{data=model.Auth}
// @Router /refresh [post]
func (controller UserController) Refresh(c *fiber.Ctx) error {
	time.Sleep(2 * time.Second)
	token := c.Locals("user").(*jwt.Token)
	jwtUser, err := controller.Auth.ParseJwtToken(token)

	if (jwtUser.UserId == uuid.UUID{} || err != nil) {
		response := model.NewSuccessResponse(new(model.Auth))
		return c.Status(200).JSON(response)
	}

	user := controller.UserService.FindLoggedUsed(c.Context())

	userModel := model.NewUserFromEntity(user)
	result := model.Auth{
		User:  &userModel,
		Token: controller.generateTokenForUser(user),
	}

	response := model.NewSuccessResponse(result)
	return c.Status(200).JSON(response)
}
