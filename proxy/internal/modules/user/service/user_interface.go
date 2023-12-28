package service

import "github.com/Bubotka/Microservices/proxy/internal/models"

type Userer interface {
	Create(in CreateIn) CreateOut
	GetByID(in GetIn) GetOut
	Update(in UpdateIn) UpdateOut
	Delete(in DeleteIn) DeleteOut
	List(in ListIn) ListOut
	Count() CountOut
}

type CreateIn struct {
	User models.User
}

type CreateOut struct {
	Message string
	Error   error
}

type GetIn struct {
	Id string
}

type GetOut struct {
	User  models.User
	Error error
}

type UpdateIn struct {
	User models.User
}

type UpdateOut struct {
	Message string
	Error   error
}

type DeleteIn struct {
	Id int
}

type DeleteOut struct {
	Message string
	Error   error
}

type ListIn struct {
	Limit  int
	Offset int
}

type ListOut struct {
	User  []models.User
	Error error
}

type CountOut struct {
	Count int
	Error error
}
