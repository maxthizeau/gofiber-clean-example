package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
	"gorm.io/gorm"
)

type questionRepository struct {
	*gorm.DB
}

func NewQuestionRepository(DB *gorm.DB) *questionRepository {
	return &questionRepository{
		DB: DB,
	}
}

func (repo *questionRepository) Create(ctx context.Context, question entity.Question) (entity.Question, error) {
	err := repo.DB.WithContext(ctx).Create(&question).Error
	exception.PanicLogging(err)

	return repo.FindById(ctx, question.Id)

}

func (repo *questionRepository) FindById(ctx context.Context, id uuid.UUID) (entity.Question, error) {
	var question entity.Question
	result := repo.DB.WithContext(ctx).Table("tb_question").
		Joins("join tb_user ON tb_question.created_by = tb_user.user_id").
		Joins("join tb_answer ON tb_answer.question_id = tb_question.question_id").
		Preload("CreatedBy").
		Preload("CorrectAnswer").
		Preload("WrongAnswers").
		Where("tb_question.question_id = ?", id).First(&question)

	if result.RowsAffected == 0 {
		return entity.Question{}, exception.NotFoundError{Message: "question not found"}
	}

	return question, nil

}
