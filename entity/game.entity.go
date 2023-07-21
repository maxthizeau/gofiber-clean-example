package entity

import (
	"time"

	"github.com/google/uuid"
)

type Game struct {
	Id        uuid.UUID `gorm:"primaryKey;column:game_id;type:varchar(36)"`
	CreatedAt time.Time
	Players   []*User     `gorm:"many2many:user_games;"`
	Questions []*Question `gorm:"many2many:game_questions;"`
}

func (Game) TableName() string {
	return "tb_game"
}
