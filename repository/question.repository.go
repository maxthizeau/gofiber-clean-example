package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/entity"
)

type QuestionRepository interface {
	Create(ctx context.Context, question entity.Question) entity.Question
	FindById(ctx context.Context, id uuid.UUID) (entity.Question, error)
}
