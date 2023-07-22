package impl

import (
	"context"

	"github.com/maxthizeau/gofiber-clean-boilerplate/common"
	"github.com/maxthizeau/gofiber-clean-boilerplate/entity"
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

func (serv *questionServiceImpl) Create(ctx context.Context, createQuestionModel model.CreateQuestionModel) model.QuestionModel {
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
	})

	wrongAnswers := []model.AnswerModel{}

	for _, a := range question.WrongAnswers {
		wrongAnswers = append(wrongAnswers, model.AnswerModel{
			Id:         a.Id,
			QuestionId: a.QuestionId,
			Label:      a.Label,
			IsCorrect:  a.IsCorrect,
		})
	}

	return model.QuestionModel{
		Id:    question.Id,
		Label: question.Label,
		CorrectAnswer: model.AnswerModel{
			Id:         question.CorrectAnswer.QuestionId,
			QuestionId: question.CorrectAnswer.QuestionId,
			IsCorrect:  question.CorrectAnswer.IsCorrect,
		},
		WrongAnswers: wrongAnswers,
	}
}
