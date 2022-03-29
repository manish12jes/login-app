package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"fmt"
	em  "github.com/labstack/echo/v4/middleware"
	middle "github.com/manish12jes/login-app/routes/middlewares"
	controller "github.com/manish12jes/login-app/controllers"
	
) 

func InitRoutes() {
	e := echo.New()
	if err := initConfig(); err != nil {
		panic(err)
	}

	e.Use(em.Logger())
	e.Use(em.Recover())
	e.Static("/", "static")

	admin := e.Group("/admin")
	admin.Use(middle.JwtAuth())
	admin.Use(middle.RefreshJwtTokens)
	admin.GET("", controller.AuthAdmin)
	admin.GET("/invoice/new", controller.NewInvoice)
	admin.GET("/invoices", controller.InvoicesIndex)


	e.GET("/user/sign_in", controller.SignInForm).Name = "userSignUpForm"
	e.POST("/user/sign_in", controller.SignIn)
	// 
	// e.Renderer = component.Renderer()
	e.Logger.Fatal(e.Start(":8080"))
}

func initConfig() error {
  viper.SetConfigFile("config/config.yaml")
  err := viper.ReadInConfig()
  if err != nil {
    fmt.Println("Config Error : ", err)
    return err
  }
  return nil
}
