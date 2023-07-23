package entity

import (
	"time"

	"github.com/google/uuid"
)

type Answer struct {
	Id         uuid.UUID `gorm:"primaryKey;column:answer_id;type:uuid;default:gen_random_uuid()"`
	Label      string    `gorm:"column:label"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	QuestionId uuid.UUID `gorm:"column:question_id;type:uuid"`
	IsCorrect  bool      `gorm:"column:is_correct"`
}

func (Answer) TableName() string {
	return "tb_answer"
}
