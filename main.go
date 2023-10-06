package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/khairulharu/restapi/internal/api"
	"github.com/khairulharu/restapi/internal/component"
	"github.com/khairulharu/restapi/internal/config"
	"github.com/khairulharu/restapi/internal/repository"
	"github.com/khairulharu/restapi/internal/service"
)

func main() {
	config := config.New()
	dbConnection := component.NewDatabase(config)

	userRepository := repository.NewUser(dbConnection)
	imageRepository := repository.NewImage(dbConnection)

	userService := service.NewUser(userRepository)
	imageService := service.NewImage(imageRepository)

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())
	api.NewApp(app, userService)
	api.NewImage(app, imageService)
	_ = app.Listen(config.Server.Host + ":" + config.Server.Port)
}
