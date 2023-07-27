package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gameRepository struct {
	*gorm.DB
}

func NewGameRepository(DB *gorm.DB) *gameRepository {
	return &gameRepository{
		DB: DB,
	}
}

func (repo *gameRepository) Create(ctx context.Context, game entity.Game) entity.Game {

	game.Id = uuid.New()
	err := repo.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&game).Error; err != nil {
			return err
		}

		return nil
	})

	exception.PanicLogging(err)

	return game
}

func (repo *gameRepository) Update(ctx context.Context, game entity.Game) entity.Game {
	err := repo.DB.WithContext(ctx).Where("game_id = ?", game.Id).Updates(&game)
	exception.PanicLogging(err)
	return game
}

func (repo *gameRepository) FindById(ctx context.Context, gameId uuid.UUID) (entity.Game, error) {
	var game entity.Game

	result := repo.DB.WithContext(ctx).
		Table("tb_game").
		Preload("Questions.Answers").
		Preload("Questions.CreatedBy").
		Preload(clause.Associations).
		Where("tb_game.game_id = ?", gameId).
		First(&game)

	if result.RowsAffected == 0 {
		return entity.Game{}, errors.New("game not found")
	}
	return game, nil
}

func (repo *gameRepository) FindByPlayerId(ctx context.Context, playerId uuid.UUID) ([]entity.Game, error) {
	var games []entity.Game

	repo.DB.WithContext(ctx).
		Model(&games).
		Preload("Players").
		Preload("Questions").
		Where("user_games.user_id = ?", playerId).
		Find(&games)

	return games, nil
}
