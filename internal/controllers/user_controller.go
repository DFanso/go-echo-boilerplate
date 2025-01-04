package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dfanso/go-echo-boilerplate/internal/models"
	"github.com/dfanso/go-echo-boilerplate/internal/services"
	"github.com/dfanso/go-echo-boilerplate/pkg/utils"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (c *UserController) GetAll(ctx echo.Context) error {
	users, err := c.service.GetAll(ctx.Request().Context())
	if err != nil {
		return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get users", err)
	}
	return utils.SuccessResponse(ctx, http.StatusOK, "Users retrieved successfully", users)
}

func (c *UserController) GetByID(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID format", err)
	}

	user, err := c.service.GetByID(ctx.Request().Context(), id)
	if err != nil {
		return utils.ErrorResponse(ctx, http.StatusNotFound, "User not found", err)
	}

	return utils.SuccessResponse(ctx, http.StatusOK, "User retrieved successfully", user)
}

func (c *UserController) Create(ctx echo.Context) error {
	var user models.User
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	fmt.Printf("Raw Request Body: %s\n", string(body))
	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))

	if err := ctx.Bind(&user); err != nil {
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
	}

	// Call BeforeCreate which includes validation
	if err := user.BeforeCreate(); err != nil {
		// Handle validation errors specifically
		if e, ok := err.(validation.Errors); ok {
			return utils.ErrorResponse(ctx, http.StatusBadRequest, "Validation failed", e)
		}
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user data", err)
	}

	if err := c.service.Create(ctx.Request().Context(), &user); err != nil {
		return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create user", err)
	}

	return utils.SuccessResponse(ctx, http.StatusCreated, "User created successfully", user)
}

func (c *UserController) Update(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID format", err)
	}

	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
	}

	user.ID = id

	// Call BeforeUpdate which includes validation
	if err := user.BeforeUpdate(); err != nil {
		// Handle validation errors specifically
		if e, ok := err.(validation.Errors); ok {
			return utils.ErrorResponse(ctx, http.StatusBadRequest, "Validation failed", e)
		}
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user data", err)
	}

	if err := c.service.Update(ctx.Request().Context(), &user); err != nil {
		return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update user", err)
	}

	return utils.SuccessResponse(ctx, http.StatusOK, "User updated successfully", user)
}

func (c *UserController) Delete(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID format", err)
	}

	if err := c.service.Delete(ctx.Request().Context(), id); err != nil {
		return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete user", err)
	}

	return utils.SuccessResponse(ctx, http.StatusOK, "User deleted successfully", nil)
}
