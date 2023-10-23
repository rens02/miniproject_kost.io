package controller

import (
	"app/config"
	"app/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateKamar(c echo.Context) error {
	// Parse JSON request body into a Kamar model
	var kamar models.Kamar
	if err := c.Bind(&kamar); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Create the new room in the database
	if err := config.DB.Create(&kamar).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create room"})
	}

	// Return a JSON response with the created room
	return c.JSON(http.StatusCreated, kamar)
}
