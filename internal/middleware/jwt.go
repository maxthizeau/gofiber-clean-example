package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth/role"
)

// Middleware JWT function
func (manager *MiddlewareManager) AuthenticateJWT(askedRoles ...role.RoleEnum) func(*fiber.Ctx) error {
	// jwtSecret := config.Get("JWT_SECRET_KEY")
	// Todo : get from config

	return jwtware.New(jwtware.Config{
		SigningKey: []byte(manager.signingKey),
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)
			jwtUser, err := manager.AuthManager.ParseJwtToken(user)

			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
					Code:    401,
					Message: "Unauthorized",
					Data:    "Undefined role",
				})
			}

			if len(askedRoles) == 0 {
				return c.Next()
			}

			// Here we could set a logger to register action made by user/role
			for _, asked := range askedRoles {
				if jwtUser.Roles.Has(asked) {
					return c.Next()
				}

			}

			return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
				Code:    401,
				Message: "Unauthorized",
				Data:    "Invalid role",
			})
		},

		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.
					Status(fiber.StatusBadRequest).
					JSON(model.GeneralResponse{
						Code:    400,
						Message: "Bad Request",
						Data:    "Missing or malformed JWT",
					})
			} else {
				return c.
					Status(fiber.StatusUnauthorized).
					JSON(model.GeneralResponse{
						Code:    401,
						Message: "Unauthorized",
						Data:    "Invalid or expired JWT",
					})
			}
		},
	})

}
