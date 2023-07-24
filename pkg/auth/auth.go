package auth

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth/role"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
)

type JwtUser struct {
	UserId uuid.UUID      `json:"user_id"`
	Roles  role.RolesList `json:"roles"`
}

type Auth interface {
	GenerateJwtToken(userId string, ttl time.Duration) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type AuthManager struct {
	signingKey  string
	ttl         time.Duration
	roleManager *role.RoleManager
}

func newJwtUser(userId uuid.UUID, roles role.RolesList) JwtUser {
	return JwtUser{
		UserId: userId,
		Roles:  roles,
	}
}

func NewAuthManager(signingKey string, ttl time.Duration) *AuthManager {
	return &AuthManager{
		signingKey:  signingKey,
		ttl:         ttl,
		roleManager: &role.RoleManager{},
	}
}

func (manager *AuthManager) ParseJwtToken(t interface{}) (JwtUser, error) {

	token, ok := t.(*jwt.Token)
	if !ok {
		return JwtUser{}, errors.New("could not parse the token")
	}

	data := token.Claims.(jwt.MapClaims)["sub"].(string)
	var jwtUser JwtUser
	err := json.Unmarshal([]byte(data), &jwtUser)

	if err != nil {
		return JwtUser{}, err
	}
	return jwtUser, nil
}

func (manager *AuthManager) GenerateJwtToken(userId uuid.UUID, roles role.RolesList) string {

	jwtUser := newJwtUser(userId, roles)
	data, err := json.Marshal(jwtUser)
	if err != nil {
		panic("error while marshaling jwt token")
	}

	claims := jwt.RegisteredClaims{
		Subject:   string(data),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(manager.ttl)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(manager.signingKey))
	exception.PanicLogging(err)

	return tokenSigned

}
