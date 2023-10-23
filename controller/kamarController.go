package controller

import (
	"app/config"
	"app/models"

	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateTipeKamar(c echo.Context) error {
	// Parse JSON request body into a TipeKamar models
	var tipeKamar models.TipeKamar
	if err := c.Bind(&tipeKamar); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Create the new room type in the database
	if err := config.DB.Create(&tipeKamar).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create room type"})
	}

	// Return a JSON response with the created room type
	return c.JSON(http.StatusCreated, tipeKamar)
}
