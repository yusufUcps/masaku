package routes

import (
	"masaku/controller"
	"masaku/middleware"
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

	return e

}
