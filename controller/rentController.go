package controller

import (
	"app/config"
	"app/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateRent(c echo.Context) error {
	// Parse JSON request body into a Rent models
	var rent models.Sewa
	if err := c.Bind(&rent); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Create the new rent in the database
	if err := config.DB.Create(&rent).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create rent"})
	}

	// Return a JSON response with the created rent
	return c.JSON(http.StatusCreated, rent)
}
