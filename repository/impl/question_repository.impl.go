package impl

import (
	"context"
	"fmt"

	"github.com/google/uuid"
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

func (repo *questionRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (entity.Question, error) {
	var question entity.Question
	result := repo.DB.WithContext(ctx).Table("tb_question").
		Joins("join tb_user ON tb_question.created_by = tb_user.user_id").
		Joins("join tb_answer ON tb_answer.question_id = tb_question.question_id").
		Preload("CreatedBy").
		Preload("CorrectAnswer").
		Preload("WrongAnswers").
		Where("tb_question.question_id = ?", id).First(&question)

	fmt.Printf("%+v", question)
	if result.RowsAffected == 0 {
		return entity.Question{}, exception.NotFoundError{Message: "question not found"}
	}

	return question, nil

}
