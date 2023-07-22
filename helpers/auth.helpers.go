package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/entity"
)

func GetJwtRoleFromUserEntity(user entity.User) []map[string]interface{} {
	var userRoles []map[string]interface{}

	for _, role := range user.UserRoles {
		userRoles = append(userRoles, map[string]interface{}{
			"role": role.Role,
		})
	}

	return userRoles
}

type JwtUser struct {
	UserId uuid.UUID
	Roles  []interface{}
}

func ParseJwtTokenFromContext(c *fiber.Ctx) (JwtUser, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	// Parse uuid
	userUuid, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return JwtUser{}, nil
	}

	return JwtUser{
		UserId: userUuid,
		Roles:  claims["roles"].([]interface{}),
	}, nil
}
