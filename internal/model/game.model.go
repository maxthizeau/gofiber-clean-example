package model

import (
	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
)

type (
	Game struct {
		Id        uuid.UUID  `json:"game_id"`
		Players   []User     `json:"players"`
		Questions []Question `json:"questions"`
	}
)

func NewGameFromEntity(gEntity entity.Game) Game {
	var q Game

	q.Id = gEntity.Id
	q.Players = NewUserArrayFromEntities(gEntity.Players)
	q.Questions = NewQuestionArrayFromEntities(gEntity.Questions)

	return q
}

func NewGameArrayFromEntities(gEntities []entity.Game) []Game {
	games := []Game{}

	for _, g := range gEntities {
		games = append(games, NewGameFromEntity(g))
	}

	return games
}
