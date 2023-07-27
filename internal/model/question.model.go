package model

import (
	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
)

type Question struct {
	Id         uuid.UUID `json:"question_id"`
	Label      string    `json:"label"`
	CreatedBy  User      `json:"created_by"`
	Answers    []Answer  `json:"answers"`
	VoteCount  int       `json:"vote_count"`
	VoteSum    int       `json:"vote_sum"`
	Difficulty int       `json:"difficulty"`
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
	q.VoteCount = qEntity.VoteCount
	q.VoteSum = qEntity.VoteSum
	q.Difficulty = qEntity.Difficulty

	return q
}

func NewQuestionArrayFromEntities(qEntities []entity.Question) []Question {
	questions := []Question{}

	for _, q := range qEntities {
		questions = append(questions, NewQuestionFromEntity(q))
	}
	return questions
}
