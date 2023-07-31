package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/middleware"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/service"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/common"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
)

func NewGameController(gameService *service.GameService, middlewareManager middleware.MiddlewareManager) *GameController {
	return &GameController{
		GameService: *gameService,
		Middlewares: middlewareManager,
	}
}

func (controller GameController) Route(app *fiber.App) {
	app.Get("/v1/api/game", controller.Middlewares.AuthenticateJWT(), controller.GetGamesForCurrentUser)
	app.Get("/v1/api/game/:id", controller.Middlewares.AuthenticateJWT(), controller.GetGame)
	app.Get("/v1/api/game/:id/result", controller.Middlewares.AuthenticateJWT(), controller.GetGameResult)
	app.Get("/v1/api/game/:id/status", controller.GetGameStatus)
	app.Post("/v1/api/game", controller.Middlewares.AuthenticateJWT(), controller.CreateGame)
	app.Post("/v1/api/game/:id/answer", controller.Middlewares.AuthenticateJWT(), controller.AnswerQuestion)
	app.Patch("/v1/api/game/:id/join", controller.Middlewares.AuthenticateJWT(), controller.JoinGame)
	app.Patch("/v1/api/game/:id/start", controller.Middlewares.AuthenticateJWT(), controller.StartGame)
}

// @Summary Get all games for current user
// @Tags game
// @Description Get all games for current user
// @Description Access restricted to: USER (as a player) or ADMIN
// @ModuleID GetGamesForCurrentUser
// @Accept  json
// @Produce  json
// @Success 200 {object} model.GeneralResponse{data=[]model.Game}
// @Router /game [get]
// @Security Bearer
func (controller GameController) GetGamesForCurrentUser(c *fiber.Ctx) error {
	games := controller.GameService.GetGamesForCurrentUser(c.Context())
	response := model.NewSuccessResponse(model.NewGameArrayFromEntities(games))
	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary Get a game by ID
// @Tags game
// @Description Get a game by ID
// @Description Access restricted to: USER
// @ModuleID GetGame
// @Accept  json
// @Produce  json
// @Param id  path  string  true "game Id"
// @Success 200 {object} model.GeneralResponse{data=model.Game}
// @Router /game/{id} [get]
// @Security Bearer
func (controller GameController) GetGame(c *fiber.Ctx) error {
	id := c.Params("id")
	game := controller.GameService.GetGame(c.Context(), uuid.MustParse(id))
	response := model.NewSuccessResponse(model.NewGameFromEntity(game))
	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary Get a game result by game ID
// @Tags game
// @Description Get a game result by game ID
// @Description Access restricted to: USER
// @ModuleID GetGameResult
// @Accept  json
// @Produce  json
// @Param id  path  string  true "game Id"
// @Success 200 {object} model.GeneralResponse{data=model.GameResult}
// @Router /game/{id}/result [get]
func (controller GameController) GetGameResult(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		exception.PanicBadRequest(errors.New("invalid game id - verify your link"))
	}
	game, userAnswers := controller.GameService.GetGameResults(c.Context(), id)
	questionResults := model.NewQuestionResultArrayFromEntities(game.Questions, userAnswers)

	response := model.NewSuccessResponse(model.NewGameResultFromEntity(game, questionResults))
	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary Get a game status by game ID
// @Tags game
// @Description Get a game status by game ID
// @Description Access not restricted
// @ModuleID GetGameStatus
// @Accept  json
// @Produce  json
// @Param id  path  string  true "game Id"
// @Success 200 {object} model.GeneralResponse{data=model.GameStatus}
// @Router /game/{id}/status [get]
func (controller GameController) GetGameStatus(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		exception.PanicBadRequest(errors.New("invalid game id - verify your link"))
	}
	gameStatus := controller.GameService.GetGameStatus(c.Context(), id)
	response := model.NewSuccessResponse(gameStatus)
	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary Create a new game
// @Tags game
// @Description User can create a game, wait for his friends to join (JoinGame), and start the game.
// @Description Access restricted to: USER
// @ModuleID CreateGame
// @Accept  json
// @Produce  json
// @Param input  body  string  true "todo: game parameters"
// @Success 200 {object} model.GeneralResponse{data=model.Game}
// @Router /game [post]
// @Security Bearer
func (controller GameController) CreateGame(c *fiber.Ctx) error {
	game := controller.GameService.NewGame(c.Context())
	response := model.NewSuccessResponse(model.NewGameFromEntity(game))
	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary Join an existing game
// @Tags game
// @Description User can join an existing game if not already a player of it, and if the game is not started.
// @Description Access restricted to: USER
// @ModuleID JoinGame
// @Accept  json
// @Produce  json
// @Param id  path  string  true "game Id"
// @Success 200 {object} model.GeneralResponse{data=model.Game}
// @Router /game/{id}/join [patch]
// @Security Bearer
func (controller GameController) JoinGame(c *fiber.Ctx) error {
	id := c.Params("id")
	game := controller.GameService.JoinGame(c.Context(), uuid.MustParse(id))
	response := model.NewSuccessResponse(model.NewGameFromEntity(game))
	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary Start an existing game
// @Tags game
// @Description One of the players can start an existing game if game is not already started.
// @Description Access restricted to: USER
// @ModuleID StartGame
// @Accept  json
// @Produce  json
// @Param id  path  string  true "game Id"
// @Success 200 {object} model.GeneralResponse{data=model.Game}
// @Router /game/{id}/start [patch]
// @Security Bearer
func (controller GameController) StartGame(c *fiber.Ctx) error {
	id := c.Params("id")
	game := controller.GameService.StartGame(c.Context(), uuid.MustParse(id))
	response := model.NewSuccessResponse(model.NewGameFromEntity(game))
	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary Answer a question in a game
// @Tags game
// @Description User can answer a question in a game if the game is started and the question is not already answered.
// @Description Access restricted to: USER
// @ModuleID AnswerQuestion
// @Accept  json
// @Produce  json
// @Param id  path  string  true "game Id"
// @Param id  path  string  true "question Id"
// @Param input  body  string  true "todo: answer info"
// @Success 200 {object} string "ok"
// @Router /game/{id}/answer [post]
// @Security Bearer
func (controller GameController) AnswerQuestion(c *fiber.Ctx) error {
	id := uuid.MustParse(c.Params("id"))
	var request model.CreateUserAnswerInput
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	common.Validate(request)

	controller.GameService.AnswerQuestionInGame(c.Context(), entity.UserAnswer{
		Text:          request.Text,
		AnswerRefer:   &request.AnswerId,
		QuestionRefer: &request.QuestionId,
		GameRefer:     &id,
	})

	response := model.NewSuccessResponse("Done")
	return c.Status(fiber.StatusOK).JSON(response)
}
