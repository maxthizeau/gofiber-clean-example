package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/helpers"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/middleware"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/service"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/common"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
)

func NewQuestionController(questionService *service.QuestionService, middlewareManager middleware.MiddlewareManager, authManager auth.AuthManager) *QuestionController {
	return &QuestionController{
		QuestionService: *questionService,
		Auth:            authManager,
		Middlewares:     middlewareManager,
		// Config:          config,
	}
}

func (controller QuestionController) Route(app *fiber.App) {
	app.Post("/v1/api/question", controller.Middlewares.AuthenticateJWT("ADMINISTRATOR"), controller.Create)
	app.Post("/v1/api/question/:id/add-answers", controller.Middlewares.AuthenticateJWT(), controller.Create)
	app.Get("/v1/api/question/:id", controller.Middlewares.AuthenticateJWT(), controller.GetQuestion)
	app.Post("/v1/api/question/:id/vote", controller.Middlewares.AuthenticateJWT(), controller.Vote)
}

func (controller QuestionController) Create(c *fiber.Ctx) error {
	var request model.CreateQuestionInput
	token := c.Locals("user")
	jwtUser, err := controller.Auth.ParseJwtToken(token)
	exception.PanicUnauthorized(err)

	err = c.BodyParser(&request)
	exception.PanicLogging(err)

	common.Validate(request)

	questionResult := controller.QuestionService.Create(c.Context(), request, jwtUser.UserId)

	response := model.NewSuccessResponse(model.NewQuestionFromEntity(questionResult))
	return c.Status(fiber.StatusOK).JSON(response)

}

func (controller QuestionController) GetQuestion(c *fiber.Ctx) error {

	id := c.Params("id")

	questionResult := controller.QuestionService.GetQuestion(c.Context(), id)
	response := model.NewSuccessResponse(model.NewQuestionFromEntity(questionResult))
	return c.Status(fiber.StatusOK).JSON(response)

}

func (controller QuestionController) Vote(c *fiber.Ctx) error {
	id := c.Params("id")

	type QuestionVoteInput struct {
		Negative bool `json:"negative"`
	}
	var request QuestionVoteInput
	err := c.BodyParser(&request)
	exception.PanicValidation(err)

	common.Validate(request)

	value := helpers.TernaryOperator[int8](request.Negative, -1, 1)

	controller.QuestionService.VoteForQuestion(c.Context(), id, value)
	return c.Status(fiber.StatusOK).JSON(model.NewSuccessResponse("Vote saved"))
}
