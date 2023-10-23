package req

import (
	"app/models"
	"app/models/web"
)

func PassBody(users web.UserRequest) *models.User {
	return &models.User{
		Name:     users.Name,
		Email:    users.Email,
		Password: users.Password,
	}
}
