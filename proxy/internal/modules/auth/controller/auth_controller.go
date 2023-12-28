package controller

import (
	"encoding/json"
	"github.com/Bubotka/Microservices/proxy/internal/infrastructure/responder"
	"github.com/Bubotka/Microservices/proxy/internal/modules/auth/service"

	"net/http"
)

type Auther interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type Auth struct {
	auth service.Auther
	responder.Responder
}

func NewAuth(auth service.Auther, responder responder.Responder) *Auth {
	return &Auth{auth: auth, Responder: responder}
}

func (a *Auth) Register(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		a.Responder.ErrorBadRequest(w, err)
		return
	}

	out := a.auth.Register(service.RegisterIn{
		Username: user.Username,
		Password: user.Password,
	})

	if out.Error != nil {
		a.Responder.ErrorBadRequest(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	a.OutputJSON(w, "Пользователь успешно зарегестрирован")
}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		a.Responder.ErrorBadRequest(w, err)
		return
	}

	out := a.auth.Login(service.LoginIn{
		Username: user.Username,
		Password: user.Password,
	})

	if out.Error != nil {
		a.Responder.ErrorUnauthorized(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	a.OutputJSON(w, out.Token)
}
