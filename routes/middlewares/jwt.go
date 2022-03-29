package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"

	"fmt"
	"time"
	"net/http"
	"github.com/manish12jes/login-app/services"
	"github.com/manish12jes/login-app/models"

)


func JwtAuth() echo.MiddlewareFunc {
	tokenKey := viper.GetString("jwt.access-token.name")
	fmt.Println("Inide jwt auth")
	return middleware.JWTWithConfig(middleware.JWTConfig {
		Claims:                  &services.Claims{},
        SigningKey:              []byte("hsf!ie42sda22)0"),
		TokenLookup:             "cookie:"+tokenKey, // "<source>:<name>"
		ErrorHandlerWithContext: services.JWTErrorHandler,
    })
}

func RefreshJwtTokens(next echo.HandlerFunc) echo.HandlerFunc {
	return func (c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*services.Claims)
		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 15*time.Minute {
			fmt.Println("Expire1 : ")
			rtc, err := c.Cookie("refresh-token")
			if err == nil && rtc != nil {
				token, err := jwt.ParseWithClaims(rtc.Value, &services.Claims{}, func(token *jwt.Token) (interface{}, error) {
					fmt.Println("Expire2 : ")
					return []byte(viper.GetString("jwt.refresh-token.secret")), nil
				})
				if err != nil {
					fmt.Println("Err : ", err)
					c.Response().Writer.WriteHeader(http.StatusUnauthorized)
				}

				if token != nil && token.Valid {
					fmt.Println("Expire3 : ")
					_ = services.GenerateTokenAndSetCookies(&models.User{Id: claims.Id, Name: claims.Name}, c )
				}
			}
		}
		return next(c)
	}
}


