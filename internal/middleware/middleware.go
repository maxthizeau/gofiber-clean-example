package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth/role"
)

type middlewareManager interface {
	AuthenticateJWT(askedRoles ...role.RoleEnum) func(*fiber.Ctx) error
}

type MiddlewareManager struct {
	*auth.AuthManager
	signingKey string
}

func NewMiddlewareManager(signingKey string, authManager *auth.AuthManager) *MiddlewareManager {
	return &MiddlewareManager{
		signingKey:  signingKey,
		AuthManager: authManager,
	}
}
