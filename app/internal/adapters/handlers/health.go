package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type healthHandler struct {
}

func NewHealthHandler() *healthHandler {
	return &healthHandler{}
}
func (hc *healthHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status":  "ok",
		"message": "I'm alive",
		"version": "1.0.0",
		"app":     "glowing-octo-memory",
	})
}
