package service

import (
	"context"

	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/helpers"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/repository"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
)

type gameService struct {
	repository.GameRepository
	repository.QuestionRepository
	auth.AuthManager
}

func NewGameService(gameRepository repository.GameRepository, questionRepository repository.QuestionRepository, authManager auth.AuthManager) *gameService {
	return &gameService{
		GameRepository:     gameRepository,
		QuestionRepository: questionRepository,
		AuthManager:        authManager,
	}
}

// Todo : add game parameters here
func (serv *gameService) NewGame(ctx context.Context) entity.Game {
	user := helpers.GetUserFromContext(ctx, &serv.AuthManager)
	questionCount := 5
	questions := serv.QuestionRepository.FindRandomQuestionsIds(ctx, questionCount)

	game := entity.Game{
		Players: []entity.User{
			{
				Id: user.UserId,
			},
		},
		Questions: questions,
	}

	game = serv.GameRepository.Create(ctx, game)

	foundGame, err := serv.GameRepository.FindById(ctx, game.Id)
	exception.PanicLogging(err)

	return foundGame
}
