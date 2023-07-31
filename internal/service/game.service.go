package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/helpers"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/repository"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth/role"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
)

type gameService struct {
	repository.GameRepository
	repository.QuestionRepository
	repository.UserAnswerRepository
	auth.AuthManager
}

func NewGameService(gameRepository repository.GameRepository, questionRepository repository.QuestionRepository, userAnswerRepository repository.UserAnswerRepository, authManager auth.AuthManager) *gameService {
	return &gameService{
		GameRepository:       gameRepository,
		QuestionRepository:   questionRepository,
		UserAnswerRepository: userAnswerRepository,
		AuthManager:          authManager,
	}
}

// Todo : add game parameters here
func (serv *gameService) NewGame(ctx context.Context) entity.Game {
	user, err := helpers.GetUserFromContext(ctx, &serv.AuthManager)
	if err != nil {
		exception.PanicUnauthorized(errors.New("user not logged in"))
	}

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

func (serv *gameService) JoinGame(ctx context.Context, gameId uuid.UUID) entity.Game {
	user, err := helpers.GetUserFromContext(ctx, &serv.AuthManager)
	if err != nil {
		exception.PanicUnauthorized(errors.New("user not logged in"))
	}

	foundGame, err := serv.GameRepository.FindById(ctx, gameId)
	exception.PanicLogging(err)

	isUserInPlayers := helpers.UserIsInArray(foundGame.Players, user.UserId)

	if foundGame.Status != entity.GameStatusCreated || isUserInPlayers {
		exception.PanicBadRequest(errors.New("game cannot be joined"))
	}

	foundGame.Players = append(foundGame.Players, entity.User{
		Id: user.UserId,
	})

	foundGame = serv.GameRepository.Update(ctx, foundGame)

	game, err := serv.GameRepository.FindById(ctx, foundGame.Id)
	exception.PanicNotFound(err)

	return game

}

func (serv *gameService) StartGame(ctx context.Context, gameId uuid.UUID) entity.Game {
	user, err := helpers.GetUserFromContext(ctx, &serv.AuthManager)
	if err != nil {
		exception.PanicUnauthorized(errors.New("user not logged in"))
	}

	foundGame, err := serv.GameRepository.FindById(ctx, gameId)
	exception.PanicLogging(err)

	isUserInPlayers := helpers.UserIsInArray(foundGame.Players, user.UserId)

	if foundGame.Status != entity.GameStatusCreated {
		exception.PanicBadRequest(errors.New("game cannot be started"))
	}

	if !isUserInPlayers {
		exception.PanicBadRequest(errors.New("you cannot start this game"))
	}

	foundGame.StartedAt = time.Now()
	foundGame.EndAt = time.Now().Add(time.Minute * 5)

	foundGame = serv.GameRepository.Update(ctx, foundGame)

	game, err := serv.GameRepository.FindById(ctx, foundGame.Id)
	exception.PanicNotFound(err)

	return game
}

func (serv *gameService) GetGameStatus(ctx context.Context, gameId uuid.UUID) model.GameStatus {
	user, userErr := helpers.GetUserFromContext(ctx, &serv.AuthManager)
	foundGame, err := serv.GameRepository.FindById(ctx, gameId)

	return model.GameStatus{
		Status:       foundGame.Status,
		IsUserInGame: userErr == nil && foundGame.HasPlayer(user.UserId),
		Exists:       err == nil,
	}
}

func (serv *gameService) GetGame(ctx context.Context, gameId uuid.UUID) entity.Game {
	user, err := helpers.GetUserFromContext(ctx, &serv.AuthManager)
	if err != nil {
		exception.PanicUnauthorized(errors.New("user not logged in"))
	}

	foundGame, err := serv.GameRepository.FindById(ctx, gameId)
	isUserInPlayers := helpers.UserIsInArray(foundGame.Players, user.UserId)

	if !isUserInPlayers && !user.Roles.Has(role.Administrator) {
		exception.PanicBadRequest(errors.New("you cannot get this game"))
	}

	exception.PanicNotFound(err)

	return foundGame
}

func (serv *gameService) GetGamesForCurrentUser(ctx context.Context) []entity.Game {
	user, err := helpers.GetUserFromContext(ctx, &serv.AuthManager)
	if err != nil {
		exception.PanicUnauthorized(errors.New("user not logged in"))
	}

	foundGames, err := serv.GameRepository.FindByPlayerId(ctx, user.UserId)

	exception.PanicNotFound(err)

	return foundGames
}

func (serv *gameService) GetGameResults(ctx context.Context, gameId uuid.UUID) (entity.Game, []entity.UserAnswer) {
	user, err := helpers.GetUserFromContext(ctx, &serv.AuthManager)
	if err != nil {
		exception.PanicUnauthorized(errors.New("user not logged in"))
	}
	foundGame, err := serv.GameRepository.FindById(ctx, gameId)
	exception.PanicNotFound(err)
	isUserInPlayers := helpers.UserIsInArray(foundGame.Players, user.UserId)

	if !isUserInPlayers && !user.Roles.Has(role.Administrator) {
		exception.PanicBadRequest(errors.New("you cannot get the result of this game"))
	}

	if foundGame.Status != entity.GameStatusEnded {
		exception.PanicBadRequest(errors.New("game is not ended"))
	}

	userAnswers, err := serv.UserAnswerRepository.FindByGameId(ctx, gameId)
	exception.PanicNotFound(err)

	return foundGame, userAnswers
}

func (serv *gameService) AnswerQuestionInGame(ctx context.Context, userAnswer entity.UserAnswer) error {
	user, err := helpers.GetUserFromContext(ctx, &serv.AuthManager)
	if err != nil {
		exception.PanicUnauthorized(errors.New("user not logged in"))
	}
	foundGame, err := serv.GameRepository.FindById(ctx, *userAnswer.GameRefer)
	exception.PanicNotFound(err)

	if !foundGame.HasPlayer(user.UserId) {
		exception.PanicBadRequest(errors.New("you cannot access this game"))
	}

	if foundGame.Status != entity.GameStatusStarted {
		exception.PanicBadRequest(errors.New("game is not started"))
	}

	// TODO : check if question is in game
	if !foundGame.HasQuestion(*userAnswer.QuestionRefer) {
		exception.PanicBadRequest(errors.New("question is not in game"))
	}

	userAnswer.Id = uuid.New()
	userAnswer.UserRefer = &user.UserId

	serv.UserAnswerRepository.Create(ctx, userAnswer)

	return nil
}
