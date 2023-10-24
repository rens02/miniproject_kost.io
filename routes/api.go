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

	//BASIC LOGIN REGISTER USER/ADMIN
	e.POST("/users/register", controller.Store)
	e.POST("/users/login", controller.LoginUser)
	e.POST("/admin/login", controller.LoginAdmin)

	//ADMIN CAN CONTROL USER
	Admin.GET("/users", controller.Index)
	Admin.GET("/users/:id", controller.Show)
	Admin.PUT("/users/:id", controller.Update)
	Admin.DELETE("/users/:id", controller.Delete)

	// ENDPOINT TIPE KAMAR (ADMIN BISA TAMBAH KURANG EDIT TIPE KAMAR, USER HANYA BISA LIHAT)
	Admin.POST("/tipe-kamar", controller.CreateTipeKamar)
	e.GET("/tipe-kamar", controller.GetAllTipeKamar)
	Admin.DELETE("/tipe-kamar/:id", controller.DeleteTipeKamar)
	Admin.PUT("/tipe-kamar/:id", controller.UpdateTipeKamar)

	// ENDPOINT KAMAR (ADMIN BISA TAMBAH KURANG EDIT KAMAR)
	Admin.POST("/kamar", controller.CreateKamar)
	Admin.GET("/kamar", controller.GetAllKamar)
	Admin.DELETE("/kamar/:id", controller.DeleteKamar)
	Admin.PUT("/kamar/:id", controller.UpdateKamar)
	e.GET("/kamar/:id", controller.GetKamarByID)

	// EDNPOINT KAMAR TERSEDIA (ADMIN BISA TAMBAH KURANG EDIT KAMAR TERSEDIA)
	Admin.POST("/kamar-tersedia", controller.CreateKamarTersedia)
	e.GET("/kamar-tersedia", controller.GetAllKamarTersedia)
	e.GET("/kamar-tersedia/:id", controller.GetKamarTersediaByID)

	return e

}
