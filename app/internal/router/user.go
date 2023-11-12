package router

import (
	"github.com/dfriveraa/glowing-octo-memory/app/internal/adapters/handlers"
	"github.com/dfriveraa/glowing-octo-memory/app/internal/adapters/middlewares"
	"github.com/dfriveraa/glowing-octo-memory/app/internal/adapters/repositories"
	"github.com/dfriveraa/glowing-octo-memory/app/internal/core/services"
	"github.com/gofiber/fiber/v2"
)

func setupUserRoutes(api fiber.Router, dbPool repositories.Db) {
	userService := services.NewUserService(dbPool)
	handler := handlers.NewUserHandler(*userService)
	api.Post("", handler.CreateNewUser)
	api.Post("login", handler.Authenticate)
	api.Get("me", middlewares.AddCurrentUser, handler.GetCurrentUser)
	api.Put("profile", middlewares.CheckRole("admin"), handler.UploadProfilePicture)
	api.Get(":userId", middlewares.CheckRole("admin"), handler.GetUserById)

}
