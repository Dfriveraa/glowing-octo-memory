package middlewares

import (
	"fmt"

	"strings"

	"github.com/dfriveraa/glowing-octo-memory/app/internal/core/services"
	"github.com/gofiber/fiber/v2"
)

var ss *services.PasswordService

func init() {
	fmt.Println("Initializing security middlewares")
	ss = services.NewPasswordService()

}

func jwtExtraction(c *fiber.Ctx) (*services.CustomClaims, error) {
	var claims *services.CustomClaims
	var err error
	token := c.Get("Authorization")
	jwtParts := strings.Split(token, " ")
	if len(jwtParts) != 2 || jwtParts[0] != "Bearer" {
		return claims, fmt.Errorf("Invalid or malformed token, must be in the format: Bearer {token}")

	}
	claims, err = ss.GetCurrentUser(jwtParts[1])
	if err != nil {
		return claims, err
	}
	c.Locals("current-user", claims)
	return claims, nil
}

func CheckRole(role string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		claims, err := jwtExtraction(c)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"detail": err.Error()})
		}
		if claims.Role != role {
			return c.Status(403).JSON(fiber.Map{
				"detail": "You are not authorized to access this resource"})
		}
		// Proceed with the next middleware
		return c.Next()
	}
}

func AddCurrentUser(c *fiber.Ctx) error {
	_, err := jwtExtraction(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"detail": err.Error()})
	}
	return c.Next()
}
