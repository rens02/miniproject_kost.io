package controller

import (
	"app/config"
	"app/helpers"
	"app/models"
	"app/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateKamarTersedia(c echo.Context) error {
	var KamarTersedia models.KamarTersedia
	if err := c.Bind(&KamarTersedia); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	KamarTersedia.Status = models.RoomStatusAvailable

	if err := config.DB.Create(&KamarTersedia).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create KamarTersedia"))
	}
	if err := config.DB.Preload("Kamar").Preload("Kamar.TipeKamar").Find(&KamarTersedia).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve KamarTersedia"))
	}
	response := helpers.ResponseAvail(KamarTersedia)

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success get KamarTersedia", response))
}

func GetAllKamarTersedia(c echo.Context) error {
	var kamarTersediaList []models.KamarTersedia
	if err := config.DB.Preload("Kamar").Preload("Kamar.TipeKamar").Where("status = ?", models.RoomStatusAvailable).Find(&kamarTersediaList).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve KamarTersedia"))
	}

	// Convert
	var responseList []helpers.KamarTersediaResponse
	for _, kamarTersedia := range kamarTersediaList {
		response := helpers.ResponseAvail(kamarTersedia)
		responseList = append(responseList, response)
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Successfully retrieved KamarTersedia", responseList))
}

func GetKamarTersediaByID(c echo.Context) error {
	id := c.Param("id")

	var KamarTersedia models.KamarTersedia
	if err := config.DB.Preload("Kamar").Preload("Kamar.TipeKamar").Where("status = ?", models.RoomStatusAvailable).First(&KamarTersedia, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve KamarTersedia"))
	}

	// Convert KamarTersedia to KamarTersediaResponse
	response := helpers.ResponseAvail(KamarTersedia)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Successfully retrieved KamarTersedia", response))
}
func EditKamarTersedia(c echo.Context) error {
	id := c.Param("id")

	var KamarTersedia models.KamarTersedia
	if err := config.DB.First(&KamarTersedia, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve KamarTersedia"))
	}

	var KamarTersediaUpdate models.KamarTersedia
	if err := c.Bind(&KamarTersediaUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	config.DB.Model(&KamarTersedia).Updates(KamarTersediaUpdate)
	if err := config.DB.Preload("Kamar").Preload("Kamar.TipeKamar").Find(&KamarTersedia).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve KamarTersedia"))
	}
	response := helpers.ResponseAvail(KamarTersedia)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Successfully updated KamarTersedia", response))
}
func DeleteKamarTersedia(c echo.Context) error {
	id := c.Param("id")

	var KamarTersedia models.KamarTersedia
	if err := config.DB.First(&KamarTersedia, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve KamarTersedia"))
	}

	if err := config.DB.Delete(&KamarTersedia).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete KamarTersedia"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Successfully deleted KamarTersedia", KamarTersedia))
}
