package controller

import (
	"encoding/json"
	"github.com/Bubotka/Microservices/proxy/internal/infrastructure/responder"
	"github.com/Bubotka/Microservices/user/domain/models"

	"github.com/Bubotka/Microservices/proxy/internal/modules/user/service"
	"github.com/go-chi/chi"

	"net/http"
)

type Userer interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByUsername(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type UserController struct {
	service service.Userer
	responder.Responder
}

func NewUserController(service service.Userer, responder responder.Responder) *UserController {
	return &UserController{service: service, Responder: responder}
}

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var user CreateRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		u.Responder.ErrorBadRequest(w, err)
		return
	}

	out := u.service.Create(service.CreateIn{User: models.User{
		ID:       0,
		Username: user.Username,
		Password: user.Password,
		IsDelete: false,
	}})

	if out.Error != nil {
		u.Responder.ErrorBadRequest(w, out.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	u.OutputJSON(w, out.Message)
}

func (u *UserController) GetByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	out := u.service.GetByUsername(service.GetIn{Username: username})
	if out.Error != nil {
		u.Responder.ErrorBadRequest(w, out.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	u.Responder.OutputJSON(w, out.User)
}

func (u *UserController) List(w http.ResponseWriter, r *http.Request) {
	out := u.service.List()

	if out.Error != nil {
		u.Responder.ErrorBadRequest(w, out.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	u.Responder.OutputJSON(w, ListResponse{
		Users: out.User,
	})
}
