package router

import (
	"github.com/dfriveraa/glowing-octo-memory/app/internal/adapters/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupHealthRoutes(api fiber.Router) {
	handler := handlers.NewHealthHandler()
	api.Get("", handler.HealthCheck)
}
