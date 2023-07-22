package service

import (
	"context"

	"github.com/maxthizeau/gofiber-clean-boilerplate/model"
)

type QuestionService interface {
	Create(ctx context.Context, createQuestionModel model.CreateQuestionModel) model.QuestionModel
}
