package controller

import (
	"app/config"
	"app/helpers"
	"app/models"
	"app/utils"
	"strconv"

	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateTipeKamar(c echo.Context) error {
	// Parse JSON request body into a TipeKamar models
	var tipeKamar models.TipeKamar
	if err := c.Bind(&tipeKamar); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
	}

	// Create the new room type in the database
	if err := config.DB.Create(&tipeKamar).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create room type"))
	}
	response := helpers.TipeKamarConvert(tipeKamar)

	// Return a JSON response with the created room type
	return c.JSON(http.StatusCreated, utils.SuccessResponse("Room type successfully created", response))
}

func GetAllTipeKamar(c echo.Context) error {
	var tipeKamars []models.TipeKamar

	if err := config.DB.Find(&tipeKamars).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve room type"))
	}
	var responseList []helpers.TipeKamarResponse
	for _, kamar := range tipeKamars {
		response := helpers.TipeKamarConvert(kamar)
		responseList = append(responseList, response)
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success", responseList))
}

func DeleteTipeKamar(c echo.Context) error {
	id := c.Param("id")

	var tipeKamar models.TipeKamar
	if err := config.DB.First(&tipeKamar, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve room type"))
	}

	if err := config.DB.Delete(&tipeKamar).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete room type"))
	}

	return c.JSON(http.StatusOK, utils.ErrorResponse("Room type successfully deleted"))
}

func UpdateTipeKamar(c echo.Context) error {
	var updatedTipeKamar models.TipeKamar

	if err := c.Bind(&updatedTipeKamar); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
	}

	TipeKamarID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid TipeKamar ID"))
	}

	// Fetch the existing TipeKamar
	existingTipeKamar := models.TipeKamar{}
	err = config.DB.First(&existingTipeKamar, TipeKamarID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("error fetching existing TipeKamar data"))
	}

	// Update TipeKamar information
	existingTipeKamar.Description = updatedTipeKamar.Description
	existingTipeKamar.Fasilitas = updatedTipeKamar.Fasilitas

	// Save the updated TipeKamar
	err = config.DB.Save(&existingTipeKamar).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to save updated"})
	}
	err = config.DB.First(&existingTipeKamar, TipeKamarID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("error fetching data"))
	}

	response := helpers.TipeKamarConvert(existingTipeKamar)
	return c.JSON(http.StatusOK, utils.SuccessResponse("Success", response))
}

func GetTipeKamarByID(c echo.Context) error {
	id := c.Param("id")

	var tipeKamar models.TipeKamar
	if err := config.DB.First(&tipeKamar, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve room type"))
	}
	response := helpers.TipeKamarConvert(tipeKamar)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success", response))
}
