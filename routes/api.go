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
	UserSecretKey := os.Getenv("USER_SECRET")

	e := echo.New()

	e.Use(middleware.NotFoundHandler)
	Admin := e.Group("")
	Admin.Use(echojwt.JWT([]byte(AdminSecretKey)))
	User := e.Group("")
	User.Use(echojwt.JWT([]byte(UserSecretKey)))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to RESTful API Services")
	})

	//BASIC LOGIN REGISTER USER/ADMIN
	e.POST("/users/register", controller.Store)
	e.POST("/users/login", controller.LoginUser)
	e.POST("/admin/login", controller.LoginAdmin)

	//ADMIN CAN CONTROL USER
	Admin.GET("/admin/users", controller.Index)
	Admin.GET("/users/:id", controller.Show)
	Admin.PUT("/admin/users/:id", controller.Update)
	Admin.DELETE("/admin/users/:id", controller.Delete)
	// User can see their profile
	User.GET("/users/profile", controller.Profile)

	// ENDPOINT TIPE KAMAR (ADMIN BISA TAMBAH KURANG EDIT TIPE KAMAR, USER HANYA BISA LIHAT)
	Admin.POST("/admin/tipe-kamar", controller.CreateTipeKamar)
	Admin.DELETE("/admin/tipe-kamar/:id", controller.DeleteTipeKamar)
	Admin.PUT("/admin/tipe-kamar/:id", controller.UpdateTipeKamar)
	e.GET("/tipe-kamar", controller.GetAllTipeKamar)

	// ENDPOINT KAMAR (ADMIN BISA TAMBAH KURANG EDIT KAMAR)
	Admin.POST("/admin/kamar", controller.CreateKamar)
	Admin.DELETE("/admin/kamar/:id", controller.DeleteKamar)
	Admin.PUT("/admin/kamar/:id", controller.UpdateKamar)
	e.GET("/kamar/:id", controller.GetKamarByID)
	e.GET("/kamar", controller.GetAllKamar)

	// EDNPOINT KAMAR TERSEDIA (ADMIN BISA TAMBAH KURANG EDIT KAMAR TERSEDIA)
	Admin.POST("/kamar-tersedia", controller.CreateKamarTersedia)
	e.GET("/kamar-tersedia", controller.GetAllKamarTersedia)
	e.GET("/kamar-tersedia/:id", controller.GetKamarTersediaByID)

	// ENDPOINT SEWA (USER BISA SEWA KAMAR)
	User.POST("/admin/sewa", controller.CreateRent)
	User.POST("/user/sewa/:rent_id/cancel", controller.CancelRent)
	User.GET("/user/sewa", controller.GetRent)
	User.GET("/user/history", controller.GetRentHistory)

	Admin.POST("/user/rekomendasi", controller.GetRecommendation)

	return e

}
