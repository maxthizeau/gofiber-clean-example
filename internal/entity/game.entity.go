package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameStatus string

const (
	GameStatusCreated GameStatus = "created"
	GameStatusStarted GameStatus = "started"
	GameStatusEnded   GameStatus = "ended"
)

type Game struct {
	Id        uuid.UUID `gorm:"primaryKey;column:game_id;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	StartedAt time.Time `gorm:"column:started_at;default:null"`
	EndAt     time.Time `gorm:"column:end_at"`
	Status    GameStatus
	Players   []User     `gorm:"many2many:user_games;"`
	Questions []Question `gorm:"many2many:game_questions;"`
}

func (Game) TableName() string {
	return "tb_game"
}

func (game *Game) AfterFind(tx *gorm.DB) (err error) {

	if game.StartedAt.IsZero() {
		game.Status = GameStatusCreated
	} else if time.Now().UTC().Before(game.EndAt) {
		game.Status = GameStatusStarted
	} else {
		game.Status = GameStatusEnded
	}

	return nil
}

func (game *Game) HasQuestion(questionId uuid.UUID) bool {
	for _, q := range game.Questions {
		if q.Id == questionId {
			return true
		}
	}
	return false
}
func (game *Game) HasPlayer(userId uuid.UUID) bool {
	for _, q := range game.Players {
		if q.Id == userId {
			return true
		}
	}
	return false
}
