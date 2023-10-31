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
	e.POST("/register", controller.Store)
	e.POST("/user/login", controller.LoginUser)
	e.POST("/admin/login", controller.LoginAdmin)

	//ADMIN CAN CONTROL USER
	Admin.GET("/admin/users", controller.Index)
	Admin.GET("/admin/user/:id", controller.Show)
	Admin.PUT("/admin/user/:id", controller.Update)
	Admin.DELETE("/admin/user/:id", controller.Delete)
	Admin.GET("/admin/user/:id_user/sewa-history", controller.GetUserHistory)
	// User can see their profile
	User.GET("/user/profile", controller.Profile)

	// ENDPOINT TIPE KAMAR (ADMIN BISA TAMBAH KURANG EDIT TIPE KAMAR, USER HANYA BISA LIHAT)
	Admin.POST("/admin/tipe-kamar", controller.CreateTipeKamar)
	Admin.DELETE("/admin/tipe-kamar/:id", controller.DeleteTipeKamar)
	Admin.PUT("/admin/tipe-kamar/:id", controller.UpdateTipeKamar)
	e.GET("/tipe-kamar", controller.GetAllTipeKamar)
	e.GET("/tipe-kamar/:id", controller.GetTipeKamarByID)

	// ENDPOINT KAMAR (ADMIN BISA TAMBAH KURANG EDIT KAMAR)
	Admin.POST("/admin/kamar", controller.CreateKamar)
	Admin.DELETE("/admin/kamar/:id", controller.DeleteKamar)
	e.GET("/kamar/:id", controller.GetKamarByID)
	e.GET("/kamar", controller.GetAllKamar)

	// EDNPOINT KAMAR TERSEDIA (ADMIN BISA TAMBAH KURANG EDIT KAMAR TERSEDIA)
	Admin.POST("/admin/kamar-tersedia", controller.CreateKamarTersedia)
	Admin.DELETE("/admin/kamar-tersedia/:id", controller.DeleteKamarTersedia)
	Admin.PUT("/admin/kamar-tersedia/:id", controller.EditKamarTersedia)
	e.GET("/kamar-tersedia", controller.GetAllKamarTersedia)
	e.GET("/kamar-tersedia/:id", controller.GetKamarTersediaByID)

	// ENDPOINT SEWA (USER BISA SEWA KAMAR)
	User.POST("/user/sewa", controller.CreateRent)
	User.DELETE("/user/sewa/:rent_id/cancel", controller.CancelRent)
	User.GET("/user/sewa", controller.GetRent)
	User.GET("/user/history", controller.GetRentHistory)

	Admin.POST("/admin/rekomendasi", controller.GetRecommendation)

	return e

}
