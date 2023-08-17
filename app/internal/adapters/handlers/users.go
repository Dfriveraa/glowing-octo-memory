package handlers

import (
	"strconv"

	"github.com/dfriveraa/glowing-octo-memory/app/internal/core/domain"
	"github.com/dfriveraa/glowing-octo-memory/app/internal/core/services"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	service services.UserService
}

func NewUserHandler(us services.UserService) *userHandler {
	return &userHandler{
		service: us,
	}
}

func (us *userHandler) CreateNewUser(c *fiber.Ctx) error {
	newUser := domain.User{Name: "John", Surname: "Doe", Email: "danielf.r97@gmail.com", Password: "123456"}
	err := us.service.CreateUser(&newUser)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Detail": "Could not create the user",
		})
	}
	return c.Status(200).JSON(newUser)
}

func (us *userHandler) GetUserById(c *fiber.Ctx) error {
	id := c.Params("userId")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Detail": "Could not validate the user id",
		})
	}
	user, err := us.service.GetUserById(userId)
	//if not user return a 404 with a message
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Detail": "User not found",
		})
	}
	return c.Status(200).JSON(user)
}
