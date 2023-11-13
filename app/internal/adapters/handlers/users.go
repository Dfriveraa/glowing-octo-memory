package handlers

import (
	"log"
	"strconv"

	"strings"

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
	registerUser := domain.RegisterUser{}
	var err error
	err = c.BodyParser(&registerUser)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"detail": "Could not parse the body",
		})
	}
	if registerUser.Password != registerUser.ConfirmPassword {
		return c.Status(400).JSON(fiber.Map{
			"detail": "Passwords do not match",
		})
	}
	newUser := domain.User{
		Name:     registerUser.Name,
		Email:    registerUser.Email,
		Password: registerUser.Password,
		Role:     registerUser.Role,
		Surname:  registerUser.Surname,
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

func (us *userHandler) GetCurrentUser(c *fiber.Ctx) error {
	a := c.Locals("current-user")
	claims := a.(*services.CustomClaims)
	profile := c.Query("profile")
	user, err := us.service.GetUserById(claims.Id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"detail": "User not found",
		})
	}

	userResponse := domain.HidePassword(user)
	if profile == "true" {
		url, err := us.service.GetProfilePicture(claims.Id)
		if err != nil {
			log.Printf("Could not get profile photo for user: %v", err)
		} else {
			userResponse.ProfilePicture = url
		}
	}
	return c.Status(200).JSON(userResponse)
}

func (us *userHandler) UploadProfilePicture(c *fiber.Ctx) error {
	a := c.Locals("current-user")
	claims := a.(*services.CustomClaims)
	userId := claims.Id
	file, err := c.FormFile("profile_picture")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"detail": "Could not parse the file",
		})
	}
	names := strings.Split(file.Filename, ".")
	err = us.service.UploadProfilePicture(userId, names[len(names)-1], file)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"detail": "Could not upload the file",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"detail": "File uploaded successfully",
	})
}

func (us *userHandler) GetProfileUrl(c *fiber.Ctx) error {
	a := c.Locals("current-user")
	claims := a.(*services.CustomClaims)
	userId := claims.Id
	url, err := us.service.GetProfilePicture(userId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"detail": "Could not get the url",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"url": url,
	})
}

func (us *userHandler) GetUserById(c *fiber.Ctx) error {
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
