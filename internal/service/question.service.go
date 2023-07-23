package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/repository"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/common"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
)

type questionService struct {
	repository.QuestionRepository
}

func NewQuestionService(questionRepository repository.QuestionRepository) *questionService {
	return &questionService{
		QuestionRepository: questionRepository,
	}
}

func (serv *questionService) Create(ctx context.Context, createQuestionModel model.CreateQuestionModel, userId uuid.UUID) entity.Question {
	common.Validate(createQuestionModel)

	wrongAnswersEntities := []entity.Answer{}

	for _, wrongLabel := range createQuestionModel.WrongAnswers {
		wrongAnswersEntities = append(wrongAnswersEntities, entity.Answer{
			Label:     wrongLabel,
			IsCorrect: false,
		})
	}

	question, err := serv.QuestionRepository.Create(ctx, entity.Question{
		Label:         createQuestionModel.Label,
		CorrectAnswer: entity.Answer{Label: createQuestionModel.CorrectAnswer, IsCorrect: true},
		WrongAnswers:  wrongAnswersEntities,
		CreatedById:   userId,
	})

	exception.PanicLogging(err)

	return question
}

func (serv *questionService) GetQuestion(ctx context.Context, id string) entity.Question {

	question, err := serv.QuestionRepository.FindById(ctx, uuid.MustParse(id))

	exception.PanicLogging(err)

	return question
}
