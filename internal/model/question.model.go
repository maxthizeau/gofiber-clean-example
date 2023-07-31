package model

import (
	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
)

type Question struct {
	Id         uuid.UUID `json:"question_id" binding:"required"`
	Label      string    `json:"label" binding:"required"`
	CreatedBy  User      `json:"created_by"`
	Answers    []Answer  `json:"answers" binding:"required"`
	VoteCount  int       `json:"vote_count" binding:"required"`
	VoteSum    int       `json:"vote_sum" binding:"required"`
	Difficulty int       `json:"difficulty" binding:"required"`
}

type QuestionResult struct {
	Question
	UserAnswers []UserAnswer `json:"user_answers"`
}

type CreateQuestionInput struct {
	Label          string   `json:"label" validate:"required"`
	CorrectAnswers []string `json:"correct_answers" validate:"required,gt=0,dive,required"`
	WrongAnswers   []string `json:"wrong_answers" validate:"required,gt=0,dive,required"`
}

type CreateUpdateAnswersForQuestionInput struct {
	Answers []CreateUpdateAnswerInput `json:"answers" validate:"required,gt=0"`
}

// Question
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

// QuestionResult
func NewQuestionResultFromEntity(qEntity entity.Question, userAnswers []entity.UserAnswer) QuestionResult {
	var q Question = NewQuestionFromEntity(qEntity)

	qResult := QuestionResult{
		Question:    q,
		UserAnswers: NewUserAnswerArrayFromEntities(userAnswers),
	}

	return qResult
}

func NewQuestionResultArrayFromEntities(qEntities []entity.Question, userAnswers []entity.UserAnswer) []QuestionResult {
	qResults := []QuestionResult{}

	for _, q := range qEntities {
		qResults = append(qResults, NewQuestionResultFromEntity(q, userAnswers))
	}

	return qResults
}
