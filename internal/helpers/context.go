package helpers

import (
	"context"

	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth"
)

func GetUserFromContext(ctx context.Context, authManager *auth.AuthManager) (auth.JwtUser, error) {
	return authManager.ParseJwtToken(ctx.Value("user"))
}
