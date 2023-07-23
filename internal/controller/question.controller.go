package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/middleware"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/service"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/common"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
)

func NewQuestionController(questionService *service.QuestionService) *QuestionController {
	return &QuestionController{
		QuestionService: *questionService,
		// Config:          config,
	}
}

func convertQuestionEntityToModel(qEntity entity.Question) model.QuestionModel {
	wrongAnswers := []model.AnswerModel{}

	for _, a := range qEntity.WrongAnswers {
		wrongAnswers = append(wrongAnswers, model.AnswerModel{
			Id:         a.Id,
			Label:      a.Label,
			QuestionId: a.QuestionId,
			IsCorrect:  a.IsCorrect,
		})
	}

	result := model.QuestionModel{
		Id:    qEntity.Id,
		Label: qEntity.Label,
		CreatedBy: model.UserModel{
			Username: qEntity.CreatedBy.Username,
			Email:    qEntity.CreatedBy.Email,
		},
		CorrectAnswer: model.AnswerModel{
			Id:         qEntity.CorrectAnswer.Id,
			Label:      qEntity.CorrectAnswer.Label,
			QuestionId: qEntity.CorrectAnswer.QuestionId,
			IsCorrect:  qEntity.CorrectAnswer.IsCorrect,
		},
		WrongAnswers: wrongAnswers,
	}

	return result
}

func (controller QuestionController) Route(app *fiber.App) {
	app.Post("/v1/api/question", middleware.AuthenticateJWT("ADMINISTRATOR"), controller.Create)
	app.Get("/v1/api/question/:id", middleware.AuthenticateJWT("ADMINISTRATOR"), controller.GetQuestion)
}

func (controller QuestionController) Create(c *fiber.Ctx) error {
	var request model.CreateQuestionModel
	token := c.Locals("user")
	jwtUser, err := auth.ParseJwtToken(token)
	exception.PanicUnauthorized(err)

	err = c.BodyParser(&request)
	exception.PanicLogging(err)

	common.Validate(request)

	questionResult := controller.QuestionService.Create(c.Context(), request, jwtUser.UserId)

	result := convertQuestionEntityToModel(questionResult)

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})

}

func (controller QuestionController) GetQuestion(c *fiber.Ctx) error {

	id := c.Params("id")

	questionResult := controller.QuestionService.GetQuestion(c.Context(), id)
	result := convertQuestionEntityToModel(questionResult)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})

}
