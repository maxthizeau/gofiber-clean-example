package common

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/configuration"
	"github.com/maxthizeau/gofiber-clean-boilerplate/exception"
)

func GenerateJwtToken(userId uuid.UUID, roles []map[string]interface{}, config configuration.Config) string {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	jwtExpired, err := strconv.Atoi(config.Get("JWT_EXPIRE_MINUTES_COUNT"))
	exception.PanicLogging(err)

	claims := jwt.MapClaims{
		"user_id": userId,
		"roles":   roles,
		"exp":     time.Now().Add(time.Minute * time.Duration(jwtExpired)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicLogging(err)

	return tokenSigned

}
