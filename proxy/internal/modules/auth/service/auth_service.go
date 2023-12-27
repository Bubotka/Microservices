package service

import (
	"fmt"
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

var users = make(map[string]string)

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) Register(in RegisterIn) RegisterOut {
	password, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return RegisterOut{
			HashedPassword: "",
			Error:          err,
		}
	}
	users[in.Username] = string(password)
	return RegisterOut{
		Error: nil,
	}
}

func (a *Auth) Login(in LoginIn) LoginOut {
	hashedPassword, ok := users[in.Username]
	if !ok {
		return LoginOut{
			Token: "",
			Error: fmt.Errorf("no such a user"),
		}
	}
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(in.Password))
	if err != nil {
		return LoginOut{
			Token: "",
			Error: err,
		}
	}

	tokenAuth := jwtauth.New("HS256", []byte("mysecretkey"), nil)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"username": in.Username})
	return LoginOut{
		Token: tokenString,
		Error: nil,
	}
}
