package handlers

import (
	"fmt"
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
	newUser := domain.User{}
	var err error
	err = c.BodyParser(&newUser)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"detail": "Could not parse the body",
		})
	}
	err = us.service.CreateUser(&newUser)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"detail": "Could not create the user",
		})
	}
	userResponse := domain.HidePassword(&newUser)
	return c.Status(200).JSON(userResponse)
}

func (us *userHandler) Authenticate(c *fiber.Ctx) error {
	user := domain.UserLogin{}
	var err error
	err = c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"detail": "Could not parse the login body",
		})
	}
	jwtToken, err := us.service.Authenticate(user.Email, user.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"detail": "Could not validate credentials",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"access_token": jwtToken,
		"token_type":   "bearer",
	})
}

func (us *userHandler) GetUserById(c *fiber.Ctx) error {
	a := c.Locals("user")
	fmt.Println(a)
	id := c.Params("userId")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"detail": "Could not validate the user id",
		})
	}
	user, err := us.service.GetUserById(userId)
	//if not user return a 404 with a message
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"detail": "User not found",
		})
	}
	userResponse := domain.HidePassword(user)
	return c.Status(200).JSON(userResponse)
}

func (us *userHandler) GetDaniel(c *fiber.Ctx) error {
	currentUser := c.Locals("current-user")
	return c.Status(200).JSON(fiber.Map{
		"detail": currentUser,
	})

}
