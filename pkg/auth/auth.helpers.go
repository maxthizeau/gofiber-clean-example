package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
)

// func GetJwtRoleFromUserEntity(user entity.User) []map[string]interface{} {
// 	var userRoles []map[string]interface{}

// 	for _, role := range user.UserRoles {
// 		userRoles = append(userRoles, map[string]interface{}{
// 			"role": role.Role,
// 		})
// 	}

// 	return userRoles
// }

type JwtUser struct {
	UserId uuid.UUID
	Roles  []string
}

func ParseJwtToken(t interface{}) (JwtUser, error) {

	token, ok := t.(*jwt.Token)
	if !ok {
		return JwtUser{}, errors.New("could not parse the token")
	}

	claims := token.Claims.(jwt.MapClaims)
	// Parse uuid
	userUuid, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return JwtUser{}, err
	}

	return JwtUser{
		UserId: userUuid,
		Roles:  claims["roles"].([]string),
	}, nil
}

func GenerateJwtToken(userId uuid.UUID, roles []string, jwtSecret string, expireTimeMinutes uint) string {

	claims := jwt.MapClaims{
		"user_id": userId,
		"roles":   roles,
		"exp":     time.Now().Add(time.Minute * time.Duration(expireTimeMinutes)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicLogging(err)

	return tokenSigned

}
