package model

import "github.com/google/uuid"

type QuestionModel struct {
	Id            uuid.UUID     `json:"id"`
	Label         string        `json:"label"`
	CreatedBy     UserModel     `json:"createdBy"`
	CorrectAnswer AnswerModel   `json:"correctAnswer"`
	WrongAnswers  []AnswerModel `json:"wrongAnswers"`
}

type CreateQuestionModel struct {
	Label         string   `json:"label"`
	CorrectAnswer string   `json:"correctAnswer"`
	WrongAnswers  []string `json:"wrongAnswers"`
}
