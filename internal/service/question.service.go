package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/helpers"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/repository"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth/role"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/common"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
)

type questionService struct {
	repository.QuestionRepository
	repository.AnswerRepository
	repository.VoteRepository
	*auth.AuthManager
}

func NewQuestionService(questionRepository repository.QuestionRepository, answerRepository repository.AnswerRepository, voteRepository repository.VoteRepository, authManager *auth.AuthManager) *questionService {
	return &questionService{
		QuestionRepository: questionRepository,
		AnswerRepository:   answerRepository,
		VoteRepository:     voteRepository,
		AuthManager:        authManager,
	}
}

func (serv *questionService) Create(ctx context.Context, createQuestionModel model.CreateQuestionInput, userId uuid.UUID) entity.Question {
	common.Validate(createQuestionModel)

	answerEntities := []entity.Answer{}

	for _, wrongLabel := range createQuestionModel.WrongAnswers {
		answerEntities = append(answerEntities, entity.Answer{
			Label:     wrongLabel,
			IsCorrect: false,
		})
	}
	for _, correctLabel := range createQuestionModel.CorrectAnswers {
		answerEntities = append(answerEntities, entity.Answer{
			Label:     correctLabel,
			IsCorrect: true,
		})
	}

	question, err := serv.QuestionRepository.Create(ctx, entity.Question{
		Label:       createQuestionModel.Label,
		Answers:     answerEntities,
		CreatedById: userId,
	})

	exception.PanicLogging(err)

	return question
}

func (serv *questionService) AddAnswer(ctx context.Context, questionId string, answerInput model.CreateUpdateAnswersForQuestionInput) entity.Question {
	userContext, err := helpers.GetUserFromContext(ctx, serv.AuthManager)
	if err != nil {
		exception.PanicUnauthorized(errors.New("user not logged in"))
	}
	question := serv.GetQuestion(ctx, questionId)

	// 1. Should be the owner or at least moderator to add answer to question
	if !userContext.Roles.Has(role.Moderator) && question.CreatedById != userContext.UserId {
		exception.PanicUnauthorized(errors.New("you are not authorized to update this question"))
	}

	answerEntities := []entity.Answer{}
	for _, a := range answerInput.Answers {
		answerEntities = append(answerEntities, entity.Answer{
			Label:      a.Label,
			IsCorrect:  a.IsCorrect,
			QuestionId: question.Id,
		})
	}

	serv.AnswerRepository.CreateMany(ctx, answerEntities)
	question = serv.GetQuestion(ctx, questionId)
	return question
}

func (serv *questionService) GetQuestion(ctx context.Context, id string) entity.Question {

	question, err := serv.QuestionRepository.FindById(ctx, uuid.MustParse(id))

	exception.PanicLogging(err)

	return question
}

func (serv *questionService) VoteForQuestion(ctx context.Context, questionId string, value int8) {
	userContext, err := helpers.GetUserFromContext(ctx, serv.AuthManager)
	if err != nil {
		exception.PanicUnauthorized(errors.New("user not logged in"))
	}
	serv.VoteRepository.Create(ctx, entity.Vote{
		Value:      value,
		QuestionId: uuid.MustParse(questionId),
		UserId:     userContext.UserId,
	})

}
