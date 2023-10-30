package controller

import (
	"app/config"
	"app/middleware"
	"app/models"
	"app/models/web"
	"app/utils"
	"app/utils/res"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Index(c echo.Context) error {
	var users []models.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve user"))
	}

	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndex(users)

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully retrieved", response))
}

func Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var updatedUser models.User

	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	var existingUser models.User
	result := config.DB.First(&existingUser, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve user"))
	}

	config.DB.Model(&existingUser).Updates(updatedUser)

	response := res.ConvertGeneral(&existingUser)

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully updated", response))
}

func Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var existingUser models.User
	result := config.DB.First(&existingUser, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve user"))
	}

	config.DB.Delete(&existingUser)

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully deleted", nil))
}

func LoginAdmin(c echo.Context) error {
	var loginRequest web.LoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	var user models.User
	if err := config.DB.Where("email = ? AND role = ?", loginRequest.Email, models.AdminRole).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Role is not Admin"))
	}

	if err := middleware.ComparePassword(user.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credentials"))
	}

	token := middleware.CreateTokenAdmin(int(user.ID), user.Name)

	// Buat respons dengan data yang diminta
	response := web.UserLoginResponse{
		Email: user.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("LoginUser successful", response))
}
