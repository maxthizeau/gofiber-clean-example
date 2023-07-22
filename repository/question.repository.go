package repository

import (
	"context"

	"github.com/maxthizeau/gofiber-clean-boilerplate/entity"
)

type QuestionRepository interface {
	Create(ctx context.Context, question entity.Question) entity.Question
}
