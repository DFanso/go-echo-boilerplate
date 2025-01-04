package main

import (
	"log"

	"github.com/dfanso/go-echo-boilerplate/config"
	"github.com/dfanso/go-echo-boilerplate/internal/controllers"
	"github.com/dfanso/go-echo-boilerplate/internal/repositories"
	"github.com/dfanso/go-echo-boilerplate/internal/routes"
	"github.com/dfanso/go-echo-boilerplate/internal/services"
	"github.com/dfanso/go-echo-boilerplate/pkg/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize MongoDB
	db, err := database.NewMongoClient(cfg.MongoDB.URI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize dependencies
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Register routes
	routes.RegisterRoutes(e, userController)

	// Start server
	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}
