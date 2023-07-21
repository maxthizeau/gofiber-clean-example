package repository

import (
	"context"

	"github.com/maxthizeau/gofiber-clean-boilerplate/entity"
)

type UserRepository interface {
	// Authentication(ctx context.Context, username string) (entity.User, error)
	Create(username string, password string, roles []string) entity.User
	FindAll(ctx context.Context) []entity.User
}
