package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/repository"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth"
)

// Todo : Remove models from services (they should only use entities)

type UserService interface {
	FindAll(ctx context.Context) (responses []entity.User)
	SignUp(ctx context.Context, authModel model.UserSignupInput) entity.User
	Authenticate(ctx context.Context, authModel model.UserLoginInput) entity.User
	FindLoggedUsed(ctx context.Context) entity.User
}

type QuestionService interface {
	Create(ctx context.Context, createQuestionModel model.CreateQuestionInput, userId uuid.UUID) entity.Question
	GetQuestion(ctx context.Context, id string) entity.Question
	VoteForQuestion(ctx context.Context, questionId string, value int8)
}

type GameService interface {
	NewGame(ctx context.Context) entity.Game
	JoinGame(ctx context.Context, gameId uuid.UUID) entity.Game
	StartGame(ctx context.Context, gameId uuid.UUID) entity.Game
	GetGame(ctx context.Context, gameId uuid.UUID) entity.Game
	GetGamesForCurrentUser(ctx context.Context) []entity.Game
	GetGameResults(ctx context.Context, gameId uuid.UUID) (entity.Game, []entity.UserAnswer)
	GetGameStatus(ctx context.Context, gameId uuid.UUID) model.GameStatus
	AnswerQuestionInGame(ctx context.Context, userAnswer entity.UserAnswer) error
}

type Services struct {
	UserService
	QuestionService
	GameService
}

type Deps struct {
	Repos *repository.Repositories
	Auth  *auth.AuthManager
}

func NewServices(deps Deps) *Services {
	return &Services{
		UserService:     NewUserService(deps.Repos.UserRepository, deps.Auth),
		QuestionService: NewQuestionService(deps.Repos.QuestionRepository, deps.Repos.AnswerRepository, deps.Repos.VoteRepository, deps.Auth),
		GameService:     NewGameService(deps.Repos.GameRepository, deps.Repos.QuestionRepository, deps.Repos.UserAnswerRepository, *deps.Auth),
	}
}
