package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
	"gorm.io/gorm"
)

type voteRepository struct {
	*gorm.DB
}

func NewVoteRepository(DB *gorm.DB) *voteRepository {
	return &voteRepository{
		DB: DB,
	}
}

func (repo *voteRepository) Create(ctx context.Context, vote entity.Vote) entity.Vote {

	vote.Id = uuid.New()

	count := repo.CountByQuestionAndUser(ctx, vote.QuestionId, vote.UserId)
	if count > 0 {
		exception.PanicBadRequest(errors.New("user already voted for this question"))
	}

	err := repo.DB.Create(&vote).Error
	exception.PanicLogging(err)

	return vote

}

func (repo *voteRepository) CountByQuestionAndUser(ctx context.Context, questionId uuid.UUID, userId uuid.UUID) (count int64) {
	err := repo.DB.WithContext(ctx).Model(&entity.Vote{}).Where("question_id = ? AND user_id = ?", questionId, userId).Count(&count).Error
	exception.PanicLogging(err)
	return count
}
