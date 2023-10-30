package controller

import (
	"app/config"
	"app/middleware"
	"app/models"
	"app/models/web"
	"app/utils"
	"app/utils/req"
	"app/utils/res"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve user"))
	}

	response := res.ConvertGeneral(&user)

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully retrieved", response))
}

func Store(c echo.Context) error {
	var user web.UserRequest
	user.Role = models.UserRole

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	userDb := req.PassBody(user)

	// Hash the user's password before storing it
	userDb.Password = middleware.HashPassword(userDb.Password)

	if err := config.DB.Create(&userDb).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to store user data"))
	}

	// Return the response without including a JWT token
	response := res.ConvertGeneral(userDb)

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", response))
}

func LoginUser(c echo.Context) error {
	var loginRequest web.LoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	var user models.User
	if err := config.DB.Where("email = ? AND role = ?", loginRequest.Email, models.UserRole).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credentials"))
	}

	if err := middleware.ComparePassword(user.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credentials"))
	}

	token := middleware.CreateTokenUser(int(user.ID), user.Name)

	// Buat respons dengan data yang diminta
	response := web.UserLoginResponse{
		Email: user.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("LoginUser successful", response))
}

func Profile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := int(claims["id"].(float64))

	var profile models.User

	if err := config.DB.First(&profile, ID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve user"))
	}

	response := res.ConvertGeneral(&profile)

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully retrieved", response))

}
