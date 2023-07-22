package impl

import (
	"context"

	"github.com/maxthizeau/gofiber-clean-boilerplate/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/exception"
	"github.com/maxthizeau/gofiber-clean-boilerplate/repository"
	"gorm.io/gorm"
)

type questionRepositoryImpl struct {
	*gorm.DB
}

func NewQuestionRepositoryImpl(DB *gorm.DB) repository.QuestionRepository {
	return &questionRepositoryImpl{
		DB: DB,
	}
}

func (repo *questionRepositoryImpl) Create(ctx context.Context, question entity.Question) entity.Question {
	err := repo.DB.WithContext(ctx).Create(&question).Error
	exception.PanicLogging(err)
	return question
}
