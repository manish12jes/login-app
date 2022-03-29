package models

import(
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id string
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func GetUser() *User {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("12345"), 8)
	return &User {
		Id: "1",
		Name: "Manish",
		Email: "manish12jes@gmail.com",
		Password: string(hashPassword),
	}
}