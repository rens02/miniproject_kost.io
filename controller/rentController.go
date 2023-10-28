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
	rent.RentStatus = models.RentStatusBooked
	if err := c.Bind(&rent); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to bind data"))
	}

	// Create the new rent in the database
	if err := config.DB.Create(&rent).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create rent"))
	}
	helpers.UpdateKamarTersediaAvailability(uint(AvailID))
	if err := config.DB.Preload("User").Preload("KamarTersedia").Preload("KamarTersedia.Kamar").Preload("KamarTersedia.Kamar.TipeKamar").Find(&rent).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to preload data"))
	}

	response := helpers.ResponseSewa(rent)
	// Return a JSON response with the created rent
	return c.JSON(http.StatusCreated, utils.SuccessResponse("Rent successfully created", response))
}

func GetRent(c echo.Context) error {
	// Extract the token from the request context
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := int(claims["id"].(float64))

	var rent []models.Sewa
	if err := config.DB.Preload("User").Preload("KamarTersedia").Preload("KamarTersedia.Kamar").Preload("KamarTersedia.Kamar.TipeKamar").Where("user_id = ? AND rent_status = ?", ID, models.RentStatusBooked).Find(&rent).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to preload data"))
	}

	var responseList []helpers.RentResponse
	for _, sewa := range rent {
		response := helpers.ResponseSewa(sewa)
		responseList = append(responseList, response)
	}

	// Return a JSON response with the created rent
	return c.JSON(http.StatusOK, utils.SuccessResponse("Rent successfully Fetched", responseList))
}

func CancelRent(c echo.Context) error {
	// Extract the token from the request context
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	UserID := int(claims["id"].(float64))

	rentID, err := strconv.Atoi(c.QueryParam("rent_id")) // Assuming the reservation ID is part of the URL
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid reservation ID"))
	}

	// Check if the reservation belongs to the requesting user
	var rent models.Sewa
	if err := config.DB.Preload("KamarTersedia").Where("id = ? AND user_id = ?", rentID, UserID).First(&rent).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Reservation not found or does not belong to the user"))
	}
	AvailableId := rent.KamarTersedia.ID
	response := helpers.ResponseSewa(rent)
	helpers.UpdateKamarTersediaAvailable(AvailableId)
	// Update the reservation status to canceled
	if err := config.DB.Model(&rent).Update("rent_status", models.RentStatusCanceled).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to cancel reservation"))
	}
	if err := config.DB.Preload("User").Preload("KamarTersedia").Preload("KamarTersedia.Kamar").Preload("KamarTersedia.Kamar.TipeKamar").Where("user_id = ?", UserID).Find(&rent).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to preload data"))
	}

	// Return a JSON response with the created rent
	return c.JSON(http.StatusOK, utils.SuccessResponse("Rent successfully canceled", response))
}

func GetRentHistory(c echo.Context) error {
	// Extract the token from the request context
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := int(claims["id"].(float64))

	var rent []models.Sewa
	if err := config.DB.Preload("User").Preload("KamarTersedia").Preload("KamarTersedia.Kamar").Preload("KamarTersedia.Kamar.TipeKamar").Where("user_id = ?", ID).Find(&rent).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to preload data"))
	}

	var responseList []helpers.RentResponse
	for _, sewa := range rent {
		response := helpers.ResponseHistory(sewa)
		responseList = append(responseList, response)
	}

	// Return a JSON response with the created rent
	return c.JSON(http.StatusOK, utils.SuccessResponse("Rent successfully Fetched", responseList))
}
