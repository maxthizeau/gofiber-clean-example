package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/entity"
)

type UserRepository interface {
	// Authentication(ctx context.Context, username string) (entity.User, error)
	Create(username string, password string, email string, roles []string) entity.User
	FindAll(ctx context.Context) []entity.User
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	FindById(ctx context.Context, id uuid.UUID) (entity.User, error)
}
