package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/middleware"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/service"
)

func NewGameController(gameService *service.GameService, middlewareManager middleware.MiddlewareManager) *GameController {
	return &GameController{
		GameService: *gameService,
		Middlewares: middlewareManager,
	}
}

func (controller GameController) Route(app *fiber.App) {
	app.Post("/v1/api/create-game", controller.Middlewares.AuthenticateJWT(), controller.CreateGame)
	app.Patch("/v1/api/join-game/:id", controller.Middlewares.AuthenticateJWT(), controller.JoinGame)
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
// @Router /create-game [post]
// @Security Bearer
func (controller GameController) CreateGame(c *fiber.Ctx) error {
	game := controller.GameService.NewGame(c.Context())
	response := model.NewSuccessResponse(model.NewGameFromEntity(game))
	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller GameController) JoinGame(c *fiber.Ctx) error {
	return nil
}
