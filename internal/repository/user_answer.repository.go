package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
	"gorm.io/gorm"
)

type userAnswerRepository struct {
	*gorm.DB
}

func NewUserAnswerRepository(DB *gorm.DB) *userAnswerRepository {
	return &userAnswerRepository{
		DB: DB,
	}
}

func (repo *userAnswerRepository) Create(ctx context.Context, userAnswer entity.UserAnswer) entity.UserAnswer {
	userAnswer.Id = uuid.New()
	err := repo.DB.WithContext(ctx).Create(&userAnswer).Error

	exception.PanicLogging(err)
	return userAnswer
}

func (repo *userAnswerRepository) FindByGameId(ctx context.Context, gameId uuid.UUID) ([]entity.UserAnswer, error) {
	var userAnswers []entity.UserAnswer

	repo.DB.WithContext(ctx).
		Table("tb_user_answer").
		Preload("Question").
		Preload("User").
		Preload("Answer").
		Where("tb_user_answer.game_refer = ?", gameId).
		Find(&userAnswers)

	return userAnswers, nil
}
func (repo *userAnswerRepository) FindByUserAndGameId(ctx context.Context, userId uuid.UUID, gameId uuid.UUID) ([]entity.UserAnswer, error) {
	var userAnswers []entity.UserAnswer

	repo.DB.WithContext(ctx).
		Table("tb_user_answer").
		Preload("Question").
		Preload("User").
		Preload("Answer").
		Where("tb_user_answer.game_refer = ? AND tb_user_answer.user_refer = ?", gameId, userId).
		Find(&userAnswers)

	return userAnswers, nil
}

// func (repo *userAnswerRepository) FindByQuestionId(ctx context.Context, questionId string) ([]entity.UserAnswer, error) {
// 	return make([]entity.UserAnswer, 0), nil
// }
