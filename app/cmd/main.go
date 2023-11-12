package main

import (
	"github.com/dfriveraa/glowing-octo-memory/app/internal/adapters/repositories"
	config "github.com/dfriveraa/glowing-octo-memory/app/internal/core"
	"github.com/dfriveraa/glowing-octo-memory/app/internal/router"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

var dbInstance *repositories.Db

func main() {
	application := fiber.New()

	dbInstance = repositories.InitDB()
	defer dbInstance.Close()
	application.Use(logger.New(logger.Config{
		Format:     "${time} ${status} ${latency} ${method} ${path}\n",
		TimeFormat: "15:04:05",
		TimeZone:   "America/New_York",
	}))

	router.SetupRoutes(application, *dbInstance)
	err := application.Listen(config.Settings.ServerPort)
	if err != nil {
		panic(err)
	}

}
