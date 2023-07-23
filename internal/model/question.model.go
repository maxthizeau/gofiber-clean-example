package model

import "github.com/google/uuid"

type QuestionModel struct {
	Id            uuid.UUID     `json:"id"`
	Label         string        `json:"label"`
	CreatedBy     UserModel     `json:"created_by"`
	CorrectAnswer AnswerModel   `json:"correct_answer"`
	WrongAnswers  []AnswerModel `json:"wrong_answers"`
}

type CreateQuestionModel struct {
	Label         string   `json:"label" validate:"required"`
	CorrectAnswer string   `json:"correct_answer" validate:"required"`
	WrongAnswers  []string `json:"wrong_answers" validate:"required,gt=0,dive,required"`
}
