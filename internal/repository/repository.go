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
}

type AnswerRepository interface {
	Create(ctx context.Context, answer entity.Answer) entity.Answer
	CreateMany(ctx context.Context, answers []entity.Answer) []entity.Answer
	FindById(ctx context.Context, id string) (entity.Answer, error)
	FindByQuestionId(ctx context.Context, questionId string) ([]entity.Answer, error)
}

type Repositories struct {
	UserRepository
	QuestionRepository
	AnswerRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository:     NewUserRepository(db),
		QuestionRepository: NewQuestionRepository(db),
		AnswerRepository:   NewAnswerRepository(db),
	}
}
