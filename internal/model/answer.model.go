package model

import (
	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
)

type (
	Answer struct {
		Id         uuid.UUID `json:"answer_id" binding:"required"`
		Label      string    `json:"label" binding:"required"`
		QuestionId uuid.UUID `json:"question_id" binding:"required"`
		IsCorrect  bool      `json:"is_correct" binding:"required"`
	}
	UserAnswer struct {
		Id         uuid.UUID `json:"user_answer_id" binding:"required"`
		Answer     Answer    `json:"answer"`
		IsCorrect  bool      `json:"is_correct"`
		Text       string    `json:"text"`
		QuestionId uuid.UUID `json:"question_id" binding:"required"`
		UserId     uuid.UUID `json:"user_id" binding:"required"`
		GameId     uuid.UUID `json:"game_id"`
	}
	CreateUpdateAnswerInput struct {
		Label     string `json:"label" validate:"required"`
		IsCorrect bool   `json:"is_correct" validate:"required"`
	}
	CreateUserAnswerInput struct {
		AnswerId   uuid.UUID `json:"answer_id"`
		Text       string    `json:"text"`
		QuestionId uuid.UUID `json:"question_id" validate:"required"`
	}
)

type AnswerModel interface {
	NewAnswerArrayFromEntities(entities []entity.Answer) []Answer
	NewAnswerFromEntity(entity entity.Answer) Answer
}

func (a Answer) ToEntity() entity.Answer {
	var entity entity.Answer
	entity.Id = a.Id
	entity.Label = a.Label
	entity.QuestionId = a.QuestionId
	entity.IsCorrect = a.IsCorrect
	return entity
}

func NewAnswerFromEntity(entity entity.Answer) Answer {
	return Answer{
		Id:         entity.Id,
		Label:      entity.Label,
		QuestionId: entity.QuestionId,
		IsCorrect:  entity.IsCorrect,
	}
}

func NewAnswerArrayFromEntities(entities []entity.Answer) []Answer {
	answers := []Answer{}
	for _, entity := range entities {

		answers = append(answers, NewAnswerFromEntity(entity))
	}
	return answers
}

// UserAnswer
func NewUserAnswerFromEntity(userAnswer entity.UserAnswer) UserAnswer {
	return UserAnswer{
		Id:         userAnswer.Id,
		Answer:     NewAnswerFromEntity(userAnswer.Answer),
		IsCorrect:  userAnswer.IsCorrect,
		Text:       userAnswer.Text,
		QuestionId: *userAnswer.QuestionRefer,
		UserId:     *userAnswer.UserRefer,
		GameId:     *userAnswer.GameRefer,
	}
}

func NewUserAnswerArrayFromEntities(userAnswers []entity.UserAnswer) []UserAnswer {
	userAnswersModel := []UserAnswer{}
	for _, userAnswer := range userAnswers {
		userAnswersModel = append(userAnswersModel, NewUserAnswerFromEntity(userAnswer))
	}
	return userAnswersModel
}
