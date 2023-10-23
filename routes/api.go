package routes

import (
	"app/controller"
	"app/middleware"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	AdminSecretKey := os.Getenv("ADMIN_SECRET")

	e := echo.New()

	e.Use(middleware.NotFoundHandler)
	Admin := e.Group("")
	Admin.Use(echojwt.JWT([]byte(AdminSecretKey)))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to RESTful API Services")
	})

	Admin.GET("/users", controller.Index)
	Admin.GET("/users/:id", controller.Show)
	e.POST("/users/register", controller.Store)
	e.POST("/users/login", controller.LoginUser)
	e.POST("/admin/login", controller.LoginAdmin)
	Admin.PUT("/users/:id", controller.Update)
	Admin.DELETE("/users/:id", controller.Delete)

	return e

}
