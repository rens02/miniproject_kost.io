package controller

import (
	"app/config"
	"app/helpers"
	"app/models"
	"app/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateKamar(c echo.Context) error {
	// Parse JSON request body into a Kamar model
	var kamar models.Kamar
	fileheader := "PhotoKamar"
	kamar.PhotoKamar = helpers.CloudinaryUpload(c, fileheader)
	if err := c.Bind(&kamar); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
	}

	// Create the new room in the database
	if err := config.DB.Create(&kamar).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create room"))
	}

	// Return a JSON response with the created room
	return c.JSON(http.StatusCreated, utils.SuccessResponse("Room successfully created", kamar))
}

func DeleteKamar(c echo.Context) error {
	id := c.Param("id")

	var kamar models.Kamar
	if err := config.DB.First(&kamar, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve room"))
	}

	if err := config.DB.Delete(&kamar).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete room"))
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Room successfully deleted"})
}

func GetAllKamar(c echo.Context) error {
	var kamars []models.Kamar

	if err := config.DB.Preload("TipeKamar").Find(&kamars).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve rooms"))
	}

	// Convert kamars to KamarResponse
	var responseList []helpers.KamarResponse
	for _, kamar := range kamars {
		response := helpers.KamarConvert(kamar)
		responseList = append(responseList, response)
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Rooms data successfully retrieved", responseList))
}

func GetKamarByID(c echo.Context) error {
	id := c.Param("id")

	var kamar models.Kamar
	if err := config.DB.Preload("TipeKamar").First(&kamar, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve room"))
	}

	// Convert kamar to KamarResponse
	response := helpers.KamarConvert(kamar)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Room data successfully retrieved", response))
}
