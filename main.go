package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/config"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/configuration"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/controller"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/repository"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/service"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/database"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/logger"
)

func main() {
	log.Println("Work in progres...")

	configPath := "configs"
	// 1. Load config
	cfg, err := config.Init(configPath)
	if err != nil {
		logger.Error(err)
		return
	}

	database := database.NewDatabase(database.DatabaseConfig{
		User:            cfg.PSQL.User,
		Password:        cfg.PSQL.Password,
		Host:            cfg.PSQL.Host,
		Port:            cfg.PSQL.Port,
		DBName:          cfg.PSQL.DBName,
		MaxPoolOpen:     cfg.PSQL.MaxPoolOpen,
		MaxPoolIdle:     cfg.PSQL.MaxPoolIdle,
		MaxPollLifeTime: cfg.PSQL.MaxPollLifeTime,
	})

	// repository
	repos := repository.NewRepositories(database)
	// service
	services := service.NewServices(service.Deps{
		Repos: repos,
	})
	// controller
	controllers := controller.NewControllers(services)

	// fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())

	// route
	controllers.ServeRoutes(app)

	err = app.Listen(cfg.HTTP.Host + ":" + cfg.HTTP.Port)
	exception.PanicLogging(err)
}
