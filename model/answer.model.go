package model

import "github.com/google/uuid"

type AnswerModel struct {
	Id         uuid.UUID `json:"id"`
	Label      string    `json:"label"`
	QuestionId uuid.UUID `json:"questionId"`
	IsCorrect  bool      `json:"isCorrect"`
}
