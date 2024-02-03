package api

import (
	"github.com/RianIhsan/go-ecommerce-app/config"
	"github.com/RianIhsan/go-ecommerce-app/internal/api/rest"
	"github.com/RianIhsan/go-ecommerce-app/internal/api/rest/handlers"
	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	rh := &rest.RestHandler{
		App: app,
	}

	setupRoutes(rh)
	app.Listen(config.ServerPort)

}

func setupRoutes(rh *rest.RestHandler) {
	// Setup user routes
	handlers.SetupUserRoutes(rh)
}
