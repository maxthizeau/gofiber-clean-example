package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
	"gorm.io/gorm"
)

type answerRepository struct {
	*gorm.DB
}

func NewAnswerRepository(DB *gorm.DB) *answerRepository {
	return &answerRepository{
		DB: DB,
	}
}

func (repo *answerRepository) Create(ctx context.Context, answer entity.Answer) entity.Answer {
	answer.Id = uuid.New()
	err := repo.DB.WithContext(ctx).Create(&answer).Error
	exception.PanicLogging(err)
	return answer
}

func (repo *answerRepository) CreateMany(ctx context.Context, answers []entity.Answer) []entity.Answer {

	for i := range answers {
		answers[i].Id = uuid.New()
	}

	err := repo.DB.WithContext(ctx).Create(&answers).Error
	exception.PanicLogging(err)
	return answers
}

func (repo *answerRepository) Update(ctx context.Context, answer entity.Answer) entity.Answer {
	err := repo.DB.WithContext(ctx).Where("answer_id = ?", answer.Id).Updates(&answer)
	exception.PanicLogging(err)
	return answer
}

func (repo *answerRepository) FindById(ctx context.Context, id string) (entity.Answer, error) {

	var answer entity.Answer

	result := repo.DB.WithContext(ctx).
		Table("tb_answer").
		Preload("Question").
		Where("answer_id = ?", id).
		First(&answer)

	if result.RowsAffected == 0 {
		return entity.Answer{}, errors.New("answer not found")
	}

	return answer, nil
}

func (repo *answerRepository) FindByQuestionId(ctx context.Context, questionId string) ([]entity.Answer, error) {
	return make([]entity.Answer, 0), nil
}
