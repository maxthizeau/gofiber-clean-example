package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxthizeau/gofiber-clean-boilerplate/configuration"
	"github.com/maxthizeau/gofiber-clean-boilerplate/exception"
	"github.com/maxthizeau/gofiber-clean-boilerplate/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/service"
)

type QuestionController struct {
	service.QuestionService
	configuration.Config
}

func NewQuestionController(questionService *service.QuestionService, config configuration.Config) *QuestionController {
	return &QuestionController{
		QuestionService: *questionService,
		Config:          config,
	}
}

func (controller QuestionController) Route(app *fiber.App) {
	app.Post("/v1/api/question", controller.Create)
}

func (controller QuestionController) Create(c *fiber.Ctx) error {
	var request model.CreateQuestionModel

	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	result := controller.QuestionService.Create(c.Context(), request)

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})

}
