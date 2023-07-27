package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuestionType string

const (
	MCQ  QuestionType = "MCQ"
	Free QuestionType = "FREE"
)

// func (ct *QuestionType) Scan(value interface{}) error {
// 	*ct = QuestionType(value.([]byte))
// 	return nil
// }

// func (ct QuestionType) Value() (driver.Value, error) {
// 	return string(ct), nil
// }

type Question struct {
	Id           uuid.UUID `gorm:"primaryKey;column:question_id;type:uuid;default:gen_random_uuid()"`
	Label        string    `gorm:"column:label"`
	Difficulty   int       `gorm:"column:difficulty;default:5"` // Difficulty between 0 to 10
	QuestionType string    `gorm:"column:question_type;default:MCQ"`
	Answers      []Answer  `gorm:"ForeignKey:QuestionId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedById  uuid.UUID `gorm:"column:created_by;type:uuid"`
	CreatedBy    User      `gorm:"ForeignKey:CreatedById"`
	Games        []Game    `gorm:"many2many:game_questions;"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	VoteCount    int
	VoteSum      int
}

func (Question) TableName() string {
	return "tb_question"
}

func (q *Question) AfterFind(tx *gorm.DB) (err error) {

	type Result struct {
		Counter uint
		Sum     int
	}

	res := Result{}

	result := tx.Table("tb_vote").Select("count(tb_vote.value) as Counter, SUM(tb_vote.value) as Sum").Where("tb_vote.question_id = ?", q.Id).Scan(&res)

	if result.Error != nil {
		return result.Error
	}

	q.VoteCount = int(res.Counter)
	q.VoteSum = res.Sum
	return nil
}
