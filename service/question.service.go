package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/model"
)

type QuestionService interface {
	Create(ctx context.Context, createQuestionModel model.CreateQuestionModel, userId uuid.UUID) entity.Question
	GetQuestion(ctx context.Context, id string) entity.Question
}
