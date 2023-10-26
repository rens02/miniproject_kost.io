package controller

import (
	"app/config"
	"app/helpers"
	"app/models"
	"app/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CreateRent(c echo.Context) error {
	// Extract the token from the request context
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := int(claims["id"].(float64))

	AvailID, _ := strconv.Atoi(c.QueryParam("available"))

	var rent models.Sewa
	rent.UserID = uint(ID)
	rent.AvailableID = uint(AvailID)
	if err := c.Bind(&rent); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Create the new rent in the database
	if err := config.DB.Create(&rent).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create rent"})
	}

	if err := config.DB.Preload("User").Preload("KamarTersedia").Preload("KamarTersedia.Kamar").Preload("KamarTersedia.Kamar.TipeKamar").Find(&rent).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to preload data"})
	}

	response := helpers.ResponseSewa(rent)
	helpers.UpdateKamarTersediaAvailability(uint(AvailID))
	// Return a JSON response with the created rent
	return c.JSON(http.StatusCreated, utils.SuccessResponse("Rent successfully created", response))
}

func GetRent(c echo.Context) error {
	// Extract the token from the request context
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := int(claims["id"].(float64))

	var rent []models.Sewa
	if err := config.DB.Preload("User").Preload("KamarTersedia").Preload("KamarTersedia.Kamar").Preload("KamarTersedia.Kamar.TipeKamar").Where("user_id = ?", ID).Find(&rent).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to preload data"})
	}

	var responseList []helpers.RentResponse
	for _, sewa := range rent {
		response := helpers.ResponseSewa(sewa)
		responseList = append(responseList, response)
	}

	// Return a JSON response with the created rent
	return c.JSON(http.StatusOK, utils.SuccessResponse("Rent successfully Fetched", responseList))
}
