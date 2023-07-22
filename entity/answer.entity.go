package entity

import (
	"time"

	"github.com/google/uuid"
)

type Answer struct {
	Id         uuid.UUID `gorm:"primaryKey;column:game_id;type:uuid;default:gen_random_uuid()"`
	CreatedAt  time.Time
	QuestionId uuid.UUID
	IsCorrect  bool
}

func (Answer) TableName() string {
	return "tb_answer"
}
