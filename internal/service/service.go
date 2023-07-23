package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/repository"
)

// Todo : Remove models from services (they should only use entities)

type UserService interface {
	FindAll(ctx context.Context) (responses []model.UserModel)
	SignUp(ctx context.Context, authModel model.UserSignupModel) entity.User
	Authenticate(ctx context.Context, authModel model.UserLoginModel) entity.User
}

type QuestionService interface {
	Create(ctx context.Context, createQuestionModel model.CreateQuestionModel, userId uuid.UUID) entity.Question
	GetQuestion(ctx context.Context, id string) entity.Question
}

type Services struct {
	UserService
	QuestionService
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	return &Services{
		UserService:     NewUserService(deps.Repos.UserRepository),
		QuestionService: NewQuestionService(deps.Repos.QuestionRepository),
	}
}
