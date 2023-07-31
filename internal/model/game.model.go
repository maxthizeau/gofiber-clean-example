package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
)

type (
	Game struct {
		Id        uuid.UUID         `json:"game_id" binding:"required"`
		Players   []User            `json:"players" binding:"required"`
		Questions []Question        `json:"questions" binding:"required"`
		Status    entity.GameStatus `json:"status" binding:"required"`
		CreatedAt time.Time         `json:"created_at" binding:"required"`
		StartedAt time.Time         `json:"started_at"`
		EndAt     time.Time         `json:"end_at"`
	}

	GameResult struct {
		Game
		QuenstionResults []QuestionResult `json:"question_results" binding:"required"`
	}

	GameStatus struct {
		Exists       bool              `json:"exists" binding:"required"`
		Status       entity.GameStatus `json:"status" `
		IsUserInGame bool              `json:"is_user_in_game" binding:"required"`
	}
)

func NewGameFromEntity(gEntity entity.Game) Game {
	var q Game

	q.Id = gEntity.Id
	q.Players = NewUserArrayFromEntities(gEntity.Players)
	q.Questions = NewQuestionArrayFromEntities(gEntity.Questions)
	q.Status = gEntity.Status
	q.CreatedAt = gEntity.CreatedAt
	q.StartedAt = gEntity.StartedAt
	q.EndAt = gEntity.EndAt

	return q
}

func NewGameArrayFromEntities(gEntities []entity.Game) []Game {
	games := []Game{}

	for _, g := range gEntities {
		games = append(games, NewGameFromEntity(g))
	}

	return games
}

func NewGameResultFromEntity(gEntity entity.Game, qResults []QuestionResult) GameResult {
	var g Game = NewGameFromEntity(gEntity)

	return GameResult{
		Game:             g,
		QuenstionResults: qResults,
	}
}
