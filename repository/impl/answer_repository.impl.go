package impl

import (
	"context"
	"errors"

	"github.com/maxthizeau/gofiber-clean-boilerplate/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/exception"
	"github.com/maxthizeau/gofiber-clean-boilerplate/repository"
	"gorm.io/gorm"
)

type answerRepositoryImpl struct {
	*gorm.DB
}

func NewAnswerRepositoryImpl(DB *gorm.DB) repository.AnswerRepositoy {
	return &answerRepositoryImpl{
		DB: DB,
	}
}

func (repo *answerRepositoryImpl) Create(ctx context.Context, answer entity.Answer) entity.Answer {
	err := repo.DB.WithContext(ctx).Create(&answer).Error
	exception.PanicLogging(err)
	return answer
}

func (repo *answerRepositoryImpl) CreateMany(ctx context.Context, answers []entity.Answer) []entity.Answer {
	err := repo.DB.WithContext(ctx).Create(&answers).Error
	exception.PanicLogging(err)
	return answers
}

func (repo *answerRepositoryImpl) FindById(ctx context.Context, id string) (entity.Answer, error) {

	var answer entity.Answer

	result := repo.DB.WithContext(ctx).
		Table("tb_answer").
		Joins("join tb_question on tb_question.question_id = tb_answer.question_id").
		Preload("Question").
		Where("answer_id = ?", id).
		First(&answer)

	if result.RowsAffected == 0 {
		return entity.Answer{}, errors.New("answer not found")
	}

	return answer, nil
}

func (repo *answerRepositoryImpl) FindByQuestionId(ctx context.Context, questionId string) ([]entity.Answer, error) {
	return make([]entity.Answer, 0), nil
}
