package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	question.Id = uuid.New()
	err := repo.DB.WithContext(ctx).Create(&question).Error
	exception.PanicLogging(err)

	return repo.FindById(ctx, question.Id)

}

func (repo *questionRepository) FindById(ctx context.Context, id uuid.UUID) (entity.Question, error) {
	var question entity.Question
	result := repo.DB.WithContext(ctx).Table("tb_question").
		Preload("CreatedBy").
		Preload("Answers").
		Where("tb_question.question_id = ?", id).First(&question)

	if result.RowsAffected == 0 {
		return entity.Question{}, exception.NotFoundError{Message: "question not found"}
	}

	return question, nil

}

func (repo *questionRepository) FindRandomQuestionsIds(ctx context.Context, count int) []entity.Question {
	var questions []entity.Question
	repo.DB.WithContext(ctx).Table("tb_question").
		Order(clause.Expr{SQL: "RAND()"}).
		Preload("CreatedBy").
		Preload("Answers").
		Limit(count).Find(&questions)

	return questions
}
