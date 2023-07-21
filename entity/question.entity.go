package entity

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	Id            uuid.UUID `gorm:"primaryKey;column:game_id;type:varchar(36)"`
	CreatedAt     time.Time
	Games         []*Game  `gorm:"many2many:game_questions;"`
	CorrectAnswer Answer   `gorm:"ForeignKey:QuestionId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	WrongAnswers  []Answer `gorm:"ForeignKey:QuestionId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Question) TableName() string {
	return "tb_question"
}
