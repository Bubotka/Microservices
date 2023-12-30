package service

import "github.com/Bubotka/Microservices/proxy/internal/models"

type Userer interface {
	Create(in CreateIn) CreateOut
	GetByUsername(in GetIn) GetOut
	List() ListOut
}

type CreateIn struct {
	User models.User
}

type CreateOut struct {
	Message string
	Error   error
}

type GetIn struct {
	Username string
}

type GetOut struct {
	User  models.User
	Error error
}

type ListOut struct {
	User  []models.User
	Error error
}
