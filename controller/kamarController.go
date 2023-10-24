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

func DeleteKamar(c echo.Context) error {
	id := c.Param("id")

	var kamar models.Kamar
	if err := config.DB.First(&kamar, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve room"})
	}

	if err := config.DB.Delete(&kamar).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete room"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Room successfully deleted"})
}

// FITUR  UPDATE JANGAN LUPA DIGANTI
func UpdateKamar(c echo.Context) error {
	id := c.Param("id")

	var kamar models.Kamar
	if err := config.DB.First(&kamar, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve room"})
	}

	if err := c.Bind(&kamar).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to bind room data"})
	}

	if err := config.DB.Save(&kamar).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update room"})
	}

	return c.JSON(http.StatusOK, kamar)
}

func GetAllKamar(c echo.Context) error {
	var kamars []models.Kamar

	if err := config.DB.Preload("TipeKamar").Find(&kamars).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve rooms"})
	}

	return c.JSON(http.StatusOK, kamars)
}

func GetKamarByID(c echo.Context) error {
	id := c.Param("id")

	var kamar models.Kamar
	if err := config.DB.Preload("TipeKamar").First(&kamar, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve room"})
	}

	return c.JSON(http.StatusOK, kamar)
}
