package model

import "github.com/google/uuid"

type AnswerModel struct {
	Id         uuid.UUID `json:"id"`
	Label      string    `json:"label"`
	QuestionId uuid.UUID `json:"question_id"`
	IsCorrect  bool      `json:"is_correct"`
}
