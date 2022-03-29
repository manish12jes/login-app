package controllers

import(
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/manish12jes/login-app/component"
	"github.com/manish12jes/login-app/models"
)

func NewInvoice(c echo.Context) error{
	data := make(map[string]interface{})
	templ, err := component.TemplateRenderer("invoice/new.html")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())

	}
	if err := templ.Execute(c.Response().Writer, data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func InvoicesIndex(c echo.Context) error {
	var customers []models.Customer

	dbHandler := component.GetDBHandler()
	dbHandler.DB().Find(&customers)
	fmt.Printf("type of customer is 1: %T\n", customers)
	fmt.Println(customers)
	templ, err := component.TemplateRenderer("invoice/list.html")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())

	}
	if err := templ.Execute(c.Response().Writer, customers); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}