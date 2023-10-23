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

func GetAllTipeKamar(c echo.Context) error {
	var tipeKamars []models.TipeKamar

	if err := config.DB.Find(&tipeKamars).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve room types"})
	}

	return c.JSON(http.StatusOK, tipeKamars)
}

func DeleteTipeKamar(c echo.Context) error {
	id := c.Param("id")

	var tipeKamar models.TipeKamar
	if err := config.DB.First(&tipeKamar, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve room type"})
	}

	if err := config.DB.Delete(&tipeKamar).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete room type"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Room type successfully deleted"})
}

func UpdateTipeKamar(c echo.Context) error {
	id := c.Param("id")

	var tipeKamar models.TipeKamar
	if err := config.DB.First(&tipeKamar, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve room type"})
	}

	if err := c.Bind(&tipeKamar).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot bind tipe kamar"})
	}

	if err := config.DB.Save(&tipeKamar).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update room type"})
	}

	return c.JSON(http.StatusOK, tipeKamar)

}
