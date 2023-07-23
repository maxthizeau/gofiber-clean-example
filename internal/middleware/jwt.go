package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/model"
)

// Middleware JWT function
func AuthenticateJWT(role string) func(*fiber.Ctx) error {
	// jwtSecret := config.Get("JWT_SECRET_KEY")
	// Todo : get from config
	jwtSecret := "SECRET"

	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			roles, ok := claims["roles"].([]interface{})

			if !ok {
				return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
					Code:    401,
					Message: "Unauthorized",
					Data:    "Undefined role",
				})
			}
			// Here we could set a logger to register action made by user/role

			for _, roleInterface := range roles {
				roleMap := roleInterface.(map[string]interface{})
				if roleMap["role"] == role {
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
