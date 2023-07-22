package entity

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	Id            uuid.UUID `gorm:"primaryKey;column:question_id;type:uuid;default:gen_random_uuid()"`
	Label         string    `gorm:"column:label"`
	CorrectAnswer Answer    `gorm:"ForeignKey:QuestionId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	WrongAnswers  []Answer  `gorm:"ForeignKey:QuestionId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedById   uuid.UUID `gorm:"column:created_by;type:uuid"`
	CreatedBy     User
	Games         []Game    `gorm:"many2many:game_questions;"`
	CreatedAt     time.Time `gorm:"column:created_at"`
}

func (Question) TableName() string {
	return "tb_question"
}
