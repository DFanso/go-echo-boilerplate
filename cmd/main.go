package main

import (
	"log"

	"github.com/dfanso/go-echo-boilerplate/config"
	"github.com/dfanso/go-echo-boilerplate/internal/controllers"
	"github.com/dfanso/go-echo-boilerplate/internal/repositories"
	"github.com/dfanso/go-echo-boilerplate/internal/routes"
	"github.com/dfanso/go-echo-boilerplate/internal/services"
	"github.com/dfanso/go-echo-boilerplate/pkg/database"

	customMiddleware "github.com/dfanso/go-echo-boilerplate/pkg/middleware"
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
	e.HideBanner = false // Show the Echo banner
	e.HidePort = false   // Show the port number

	// Middleware
	e.Use(customMiddleware.NewCustomLogger().Middleware())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize dependencies
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Register routes
	routes.RegisterRoutes(e, userController)

	// health check route
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status": "OK",
		})
	})

	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}
