package router

import (
	"github.com/dfriveraa/glowing-octo-memory/app/internal/adapters/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, dbPool repositories.Db) {
	api := app.Group("/api/v1")
	healthcheck := api.Group("/healthcheck")
	setupHealthRoutes(healthcheck)
	users := api.Group("/users")
	setupUserRoutes(users, dbPool)

}
