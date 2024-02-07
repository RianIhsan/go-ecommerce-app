package api

import (
	"log"

	"github.com/RianIhsan/go-ecommerce-app/config"
	"github.com/RianIhsan/go-ecommerce-app/internal/api/rest"
	"github.com/RianIhsan/go-ecommerce-app/internal/api/rest/handlers"
	"github.com/RianIhsan/go-ecommerce-app/internal/domain"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()
	db , err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("database not connected: %v\n", err)
	}

	log.Println("database connected")

	// migration
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatal("migration failed")
	}
	log.Println("migration success")

	rh := &rest.RestHandler{
		App: app,
		DB: db,
	}

	setupRoutes(rh)
	app.Listen(config.ServerPort)

}

func setupRoutes(rh *rest.RestHandler) {
	// Setup user routes
	handlers.SetupUserRoutes(rh)
}
