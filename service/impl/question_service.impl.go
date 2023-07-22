package impl

import (
	"context"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/common"
	"github.com/maxthizeau/gofiber-clean-boilerplate/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/exception"
	"github.com/maxthizeau/gofiber-clean-boilerplate/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/repository"
	"github.com/maxthizeau/gofiber-clean-boilerplate/service"
)

type questionServiceImpl struct {
	repository.QuestionRepository
}

func NewQuestionServiceImpl(questionRepository *repository.QuestionRepository) service.QuestionService {
	return &questionServiceImpl{
		QuestionRepository: *questionRepository,
	}
}

func (serv *questionServiceImpl) Create(ctx context.Context, createQuestionModel model.CreateQuestionModel, userId uuid.UUID) entity.Question {
	common.Validate(createQuestionModel)

	wrongAnswersEntities := []entity.Answer{}

	for _, wrongLabel := range createQuestionModel.WrongAnswers {
		wrongAnswersEntities = append(wrongAnswersEntities, entity.Answer{
			Label:     wrongLabel,
			IsCorrect: false,
		})
	}

	question := serv.QuestionRepository.Create(ctx, entity.Question{
		Label:         createQuestionModel.Label,
		CorrectAnswer: entity.Answer{Label: createQuestionModel.CorrectAnswer, IsCorrect: true},
		WrongAnswers:  wrongAnswersEntities,
		CreatedById:   userId,
	})

	return question
}

func (serv *questionServiceImpl) GetQuestion(ctx context.Context, id string) entity.Question {

	question, err := serv.QuestionRepository.FindById(ctx, uuid.MustParse(id))

	exception.PanicLogging(err)

	return question
}
