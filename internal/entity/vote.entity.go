package entity

import (
	"github.com/google/uuid"
)

type Vote struct {
	Id         uuid.UUID `gorm:"primaryKey;column:vote_id;type:uuid;default:gen_random_uuid()"`
	Value      int8      `gorm:"column:value;check:value >= -1 AND value <= 1"`
	QuestionId uuid.UUID `gorm:"column:question_id;index:idx_uservote,unique"`
	UserId     uuid.UUID `gorm:"column:user_id;index:idx_uservote,unique"`
}

func (Vote) TableName() string {
	return "tb_vote"
}
