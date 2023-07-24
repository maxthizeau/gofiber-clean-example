package helpers

import (
	"context"

	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth"
)

func GetUserFromContext(ctx context.Context, authManager *auth.AuthManager) auth.JwtUser {
	user, err := authManager.ParseJwtToken(ctx.Value("user"))
	if err != nil {
		panic(err)
	}
	return user
}
