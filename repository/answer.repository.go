package repository

import (
	"context"

	"github.com/maxthizeau/gofiber-clean-boilerplate/entity"
)

type AnswerRepositoy interface {
	Create(ctx context.Context, answer entity.Answer) entity.Answer
	CreateMany(ctx context.Context, answers []entity.Answer) []entity.Answer
	FindById(ctx context.Context, id string) (entity.Answer, error)
	FindByQuestionId(ctx context.Context, questionId string) ([]entity.Answer, error)
}
