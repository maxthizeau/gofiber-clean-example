package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserAnswer struct {
	Id            uuid.UUID `gorm:"primaryKey;column:user_answer_id;type:uuid;default:gen_random_uuid()"`
	Text          string    `gorm:"column:text;default:null"` // Can have either a text (free answer) or an answer attached (MCQ)
	AnswerRefer   *uuid.UUID
	Answer        Answer `gorm:"ForeignKey:AnswerRefer;"`
	UserRefer     *uuid.UUID
	User          *User `gorm:"ForeignKey:UserRefer;constraint:OnUpdate:CASCADE,ONDELETE:SET NULL;"`
	QuestionRefer *uuid.UUID
	Question      *Question `gorm:"ForeignKey:QuestionRefer;"`
	GameRefer     *uuid.UUID
	Game          *Game     `gorm:"ForeignKey:GameRefer;"`
	IsCorrect     bool      `gorm:"column:is_correct;default:null"` // null = correction not done yet
	CreatedAt     time.Time `gorm:"column:created_at"`
}

func (UserAnswer) TableName() string {
	return "tb_user_answer"
}
