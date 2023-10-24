package controller

import (
	"app/config"
	"app/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateKamarTersedia(c echo.Context) error {

	// Parse JSON request body into the KamarTersedia model
	var KamarTersedia models.KamarTersedia
	if err := c.Bind(&KamarTersedia); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	KamarTersedia.Status = models.RoomStatusAvailable

	// Create the new KamarTersedia in the database
	if err := config.DB.Create(&KamarTersedia).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create KamarTersedia"})
	}

	// Return a JSON response with the created KamarTersedia
	return c.JSON(http.StatusCreated, KamarTersedia)
}

func GetAllKamarTersedia(c echo.Context) error {
	var kamarTersediaList []models.KamarTersedia
	if err := config.DB.Preload("Kamar").Preload("Kamar.TipeKamar").Where("status = ?", models.RoomStatusAvailable).Find(&kamarTersediaList).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve KamarTersedia"})
	}

	// Convert KamarTersedia items to KamarTersediaResponse items
	var responseList []models.KamarTersediaResponse
	for _, kamarTersedia := range kamarTersediaList {
		response := models.ResponseAvail(kamarTersedia)
		responseList = append(responseList, response)
	}

	return c.JSON(http.StatusOK, responseList)
}

func GetKamarTersediaByID(c echo.Context) error {
	id := c.Param("id")

	var KamarTersedia models.KamarTersedia
	if err := config.DB.Preload("Kamar").Preload("Kamar.TipeKamar").Where("status = ?", models.RoomStatusAvailable).First(&KamarTersedia, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve KamarTersedia"})
	}

	// Convert KamarTersedia to KamarTersediaResponse
	response := models.ResponseAvail(KamarTersedia)

	return c.JSON(http.StatusOK, response)
}
