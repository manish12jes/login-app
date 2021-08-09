package routes

import (
	"github.com/labstack/echo/v4"
	controller "github.com/manish12jes/login-app/controllers"
) 

func initRoutes() {
	e = echo.New()

	adminGroup := e.Group("admin")
	adminGroup.GET("", controller.authAdmin())
	// 
	e.Logger.Fatal(e.Start(":8080"))
}