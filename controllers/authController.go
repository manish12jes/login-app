package controllers

import(
	"fmt"
	"net/http"
	"html/template"
	"path"
	"github.com/labstack/echo/v4"
	model "github.com/manish12jes/login-app/models"
	service "github.com/manish12jes/login-app/services"
	"github.com/manish12jes/login-app/component"
)
func AuthAdmin(c echo.Context) error {
	data := make(map[string]interface{})
	templ, err := component.TemplateRenderer("index.html")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())

	}
	if err := templ.Execute(c.Response().Writer, data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func AuthSome(c echo.Context) error {
	fmt.Println("Hi, Inside someAdmin controller")
	c.String(http.StatusOK, "Hi, Inside SomeAdmin controller")
	return nil
}

func SignInForm(c echo.Context) error{
	fp := path.Join("view", "sign_in/signIn.html")
	templ, err := template.ParseFiles(fp)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())

	}
	if err := templ.Execute(c.Response().Writer, nil); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func SignIn(c echo.Context) error {
	params := &service.SignInRequest{}
	if err :=  c.Bind(params); err != nil {
		fmt.Println(err.Error())
		echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		return nil
	}
	user := model.GetUser()
	err := params.SignInUser(user, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.Redirect(http.StatusFound, "/admin")
}