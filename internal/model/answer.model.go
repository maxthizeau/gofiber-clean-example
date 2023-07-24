package model

import (
	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
)

type (
	Answer struct {
		Id         uuid.UUID `json:"answer_id"`
		Label      string    `json:"label"`
		QuestionId uuid.UUID `json:"question_id"`
		IsCorrect  bool      `json:"is_correct"`
	}
	CreateUpdateAnswerInput struct {
		Label     string `json:"label" validate:"required"`
		IsCorrect bool   `json:"is_correct" validate:"required"`
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
