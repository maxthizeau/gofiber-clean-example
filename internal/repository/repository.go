package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(username string, password string, email string, roles []string) entity.User
	FindAll(ctx context.Context) []entity.User
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	FindById(ctx context.Context, id uuid.UUID) (entity.User, error)
}

type QuestionRepository interface {
	Create(ctx context.Context, question entity.Question) (entity.Question, error)
	FindById(ctx context.Context, id uuid.UUID) (entity.Question, error)
	FindRandomQuestionsIds(ctx context.Context, count int) []entity.Question
}

type AnswerRepository interface {
	Create(ctx context.Context, answer entity.Answer) entity.Answer
	CreateMany(ctx context.Context, answers []entity.Answer) []entity.Answer
	FindById(ctx context.Context, id string) (entity.Answer, error)
	FindByQuestionId(ctx context.Context, questionId string) ([]entity.Answer, error)
}
type GameRepository interface {
	Create(ctx context.Context, game entity.Game) entity.Game
	Update(ctx context.Context, game entity.Game) entity.Game
	FindById(ctx context.Context, gameId uuid.UUID) (entity.Game, error)
	FindByPlayerId(ctx context.Context, playerId uuid.UUID) ([]entity.Game, error)
}
type VoteRepository interface {
	Create(ctx context.Context, vote entity.Vote) entity.Vote
}

type UserAnswerRepository interface {
	Create(ctx context.Context, userAnswer entity.UserAnswer) entity.UserAnswer
	FindByGameId(ctx context.Context, gameId uuid.UUID) ([]entity.UserAnswer, error)
	FindByUserAndGameId(ctx context.Context, userId uuid.UUID, gameId uuid.UUID) ([]entity.UserAnswer, error)
}

type Repositories struct {
	UserRepository
	QuestionRepository
	AnswerRepository
	GameRepository
	VoteRepository
	UserAnswerRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository:       NewUserRepository(db),
		QuestionRepository:   NewQuestionRepository(db),
		AnswerRepository:     NewAnswerRepository(db),
		GameRepository:       NewGameRepository(db),
		VoteRepository:       NewVoteRepository(db),
		UserAnswerRepository: NewUserAnswerRepository(db),
	}
}
