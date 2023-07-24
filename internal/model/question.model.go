package model

import (
	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
)

type Question struct {
	Id        uuid.UUID `json:"question_id"`
	Label     string    `json:"label"`
	CreatedBy User      `json:"created_by"`
	Answers   []Answer  `json:"answers"`
}

type CreateQuestionInput struct {
	Label          string   `json:"label" validate:"required"`
	CorrectAnswers []string `json:"correct_answers" validate:"required,gt=0,dive,required"`
	WrongAnswers   []string `json:"wrong_answers" validate:"required,gt=0,dive,required"`
}

type CreateUpdateAnswersForQuestionInput struct {
	Answers []CreateUpdateAnswerInput `json:"answers" validate:"required,gt=0"`
}

func NewQuestionFromEntity(qEntity entity.Question) Question {
	var q Question

	q.Id = qEntity.Id
	q.Label = qEntity.Label
	q.CreatedBy = NewUserFromEntity(qEntity.CreatedBy)
	q.Answers = NewAnswerArrayFromEntities(qEntity.Answers)

	return q
}
