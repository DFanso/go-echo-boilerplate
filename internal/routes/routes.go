package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/dfanso/go-echo-boilerplate/internal/controllers"
)

// RegisterRoutes registers all application routes
func RegisterRoutes(e *echo.Echo, userController *controllers.UserController) {
	// API group
	api := e.Group("/api")

	// Register all routes
	registerUserRoutes(api, userController)
}

// registerUserRoutes registers user-related routes
func registerUserRoutes(api *echo.Group, userController *controllers.UserController) {

	health := api.Group("/health")
	{
		health.GET("", func(c echo.Context) error {
			return c.JSON(200, map[string]string{"status": "OK"})
		})
	}
	users := api.Group("/users")
	{
		users.GET("", userController.GetAll)
		users.GET("/:id", userController.GetByID)
		users.POST("", userController.Create)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Delete)
	}
}