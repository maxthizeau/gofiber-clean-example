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
	app.Put("/v1/api/question/:id/vote", controller.Middlewares.AuthenticateJWT(), controller.Vote)
}

// @Summary Create a question
// @Tags question
// @Description Create a new question, associated to the current user
// @Description Access restricted to: ADMIN
// @ModuleID Create
// @Accept  json
// @Produce  json
// @Param input body model.CreateQuestionInput true "question and answers info"
// @Success 200 {object} model.GeneralResponse{data=model.Question}
// @Failure 400 {object} model.GeneralResponse{data=[]common.ValidationResponse}
// @Router /question [post]
// @Security Bearer
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

// @Summary Get a question by ID
// @Tags question
// @Description Get a question by ID
// @Description Access restricted to: USER
// @ModuleID GetQuestion
// @Accept  json
// @Produce  json
// @Param id path string true "question id"
// @Success 200 {object} model.GeneralResponse{data=model.Question}
// @Router /question/{id} [get]
// @Security Bearer
func (controller QuestionController) GetQuestion(c *fiber.Ctx) error {

	id := c.Params("id")

	questionResult := controller.QuestionService.GetQuestion(c.Context(), id)
	response := model.NewSuccessResponse(model.NewQuestionFromEntity(questionResult))
	return c.Status(fiber.StatusOK).JSON(response)

}

// @Summary User vote for a question
// @Tags question
// @Description User can upvote/downvote a question. Downvote are done by passing the "negative" prop in the request body.
// @Description It will upsert the vote -> Creating if user/question relation does not exist, update it otherwise.
// @Description Access restricted to: USER
// @ModuleID Vote
// @Accept  json
// @Produce  json
// @Param id  path  string  true "question id"
// @Param input  body  questionVoteInput  true "question vote info"
// @Success 200 {object} model.GeneralResponse{data=string}
// @Router /question/{id}/vote [put]
// @Security Bearer
func (controller QuestionController) Vote(c *fiber.Ctx) error {
	id := c.Params("id")

	var request questionVoteInput
	err := c.BodyParser(&request)
	exception.PanicValidation(err)

	common.Validate(request)

	value := helpers.TernaryOperator[int8](request.Negative, -1, 1)

	controller.QuestionService.VoteForQuestion(c.Context(), id, value)
	return c.Status(fiber.StatusOK).JSON(model.NewSuccessResponse("Vote saved"))
}

type questionVoteInput struct {
	Negative bool `json:"negative"`
}
