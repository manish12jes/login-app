package services

import (
	"errors"
	"time"
	"net/http"
	// "fmt"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	model "github.com/manish12jes/login-app/models"
)


type SignInRequest struct {
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Claims struct {
	Id string
	Name string
	jwt.StandardClaims
}

func (params SignInRequest) SignInUser(user *model.User, c echo.Context) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		return errors.New("Invalid password")
	}

	if err := GenerateTokenAndSetCookies(user, c); err != nil {
		return err
	}

	return nil
}

func GenerateTokenAndSetCookies(user *model.User, c echo.Context) error{
	accessToken, expireAt, err := generateAccessToken(user)
	if err != nil {
		return err
	}
	setTokenCookies(viper.GetString("jwt.access-token.name"), accessToken, expireAt, c)

	refreshToken, expireAt, err := generateRefreshToken(user)
	if err != nil {
		return err
	}
	setTokenCookies(viper.GetString("jwt.refresh-token.name"), refreshToken, expireAt, c)

	setUserCookies(user, c)
	return nil
}

func generateToken(user *model.User, expireTime time.Time, secret string) (string, time.Time, error){
	claims := Claims {
		Id: user.Id,
		Name: user.Name,
		StandardClaims: jwt.StandardClaims {
			ExpiresAt: expireTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", time.Now(), err
	}
	return  jwtToken, expireTime, err
}

func generateRefreshToken(user *model.User) (string, time.Time, error){
	expireTime := time.Now().Add(3 * time.Hour)
	secretKey := viper.GetString("jwt.refresh-token.secret")
	return generateToken(user, expireTime, secretKey)
}

func generateAccessToken(user *model.User) (string, time.Time, error){
	expireTime := time.Now().Add(30 * time.Minute)
	secretKey := viper.GetString("jwt.access-token.secret")
	return generateToken(user, expireTime, secretKey)
}

func setTokenCookies(tokenName string, token string, expireAt time.Time, c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = tokenName
	cookie.Value = token
	cookie.Expires = expireAt
	cookie.Path = "/"
	cookie.HttpOnly = true

	c.SetCookie(cookie)
	return nil
}

func setUserCookies(user *model.User, c echo.Context) error {
	expireTime := time.Now().Add(30 * time.Minute)

	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = user.Name
	cookie.Expires = expireTime
	cookie.Path = "/"
	cookie.HttpOnly = true

	c.SetCookie(cookie)
	return nil
}

func JWTErrorHandler(err error, c echo.Context) error {
	// fmt.Println("Error fiund", err, c.Echo().Reverse("userSignUpForm"))
	return c.Redirect(http.StatusTemporaryRedirect, c.Echo().Reverse("userSignUpForm"))
}





