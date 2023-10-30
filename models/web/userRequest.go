package web

import "app/models"

type UserRequest struct {
	Name     string      `json:"name" form:"name"`
	Email    string      `json:"email" form:"email"`
	Password string      `json:"password" form:"password"`
	Role     models.Role `json:"role" form:"role"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
