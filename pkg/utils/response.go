package utils

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(c echo.Context, code int, message string, data interface{}) error {
	return c.JSON(code, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c echo.Context, code int, message string, err error) error {
	return c.JSON(code, Response{
		Status:  "error",
		Message: message,
		Error:   err.Error(),
	})
}
