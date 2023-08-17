package main

import (
	"github.com/dfriveraa/glowing-octo-memory/app/internal/adapters/repositories"
	"github.com/dfriveraa/glowing-octo-memory/app/internal/router"

	"github.com/gofiber/fiber/v2"
)

var dbInstance *repositories.Db

func main() {
	application := fiber.New()

	dbInstance = repositories.InitDB()
	defer dbInstance.Close()
	
	router.SetupRoutes(application, *dbInstance)
	err := application.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
