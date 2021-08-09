package controllers

import(
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)
func authAdmin(e echo.Context) error {
	fmt.Println("Hi, Inside authAdmin controller")
	c.String(http.StatusOK, "Hi, Inside authAdmin controller")
	return nil
}