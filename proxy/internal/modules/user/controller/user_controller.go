package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"main/internal/infrastructure/responder"
	"main/internal/models"
	"main/internal/modules/user/service"
	"net/http"
	"strconv"
)

type Userer interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByUsername(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type User struct {
	service service.Userer
	responder.Responder
}

func NewUser(service service.Userer, responder responder.Responder) *User {
	return &User{service: service, Responder: responder}
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
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

func (u *User) GetByUsername(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println(id)
	out := u.service.GetByID(service.GetIn{Id: id})
	if out.Error != nil {
		u.Responder.ErrorBadRequest(w, out.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	u.Responder.OutputJSON(w, out.User)
}

func (u *User) Update(w http.ResponseWriter, r *http.Request) {
	var user UpdateRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		u.Responder.ErrorBadRequest(w, err)
		return
	}

	out := u.service.Update(service.UpdateIn{User: models.User{
		ID:       user.Id,
		Username: user.Username,
		Password: user.Password,
		IsDelete: false,
	}})

	if out.Error != nil {
		u.Responder.ErrorBadRequest(w, out.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	u.Responder.OutputJSON(w, out.Message)
}

func (u *User) Delete(w http.ResponseWriter, r *http.Request) {
	var user DeleteRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		u.Responder.ErrorBadRequest(w, err)
		return
	}

	out := u.service.Delete(service.DeleteIn{Id: user.Id})

	if out.Error != nil {
		u.Responder.ErrorBadRequest(w, out.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	u.Responder.OutputJSON(w, out.Message)
}

func (u *User) List(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("Limit"))
	if err != nil {
		limit = 100
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("Offset"))
	if err != nil {
		offset = 0
	}

	out := u.service.List(service.ListIn{
		Limit:  limit,
		Offset: offset,
	})

	if out.Error != nil {
		u.Responder.ErrorBadRequest(w, out.Error)
		return
	}

	outCount := u.service.Count()

	if outCount.Error != nil {
		u.Responder.ErrorBadRequest(w, out.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	u.Responder.OutputJSON(w, ListResponse{
		Users: out.User,
		Count: outCount.Count,
	})

}
