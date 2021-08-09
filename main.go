package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/manish12jes/login-app/routes"
)

func main() {
	
	routes.initRoutes()
}