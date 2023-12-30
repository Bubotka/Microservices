package service

import (
	"github.com/Bubotka/Microservices/proxy/internal/models"
	"github.com/Bubotka/Microservices/proxy/pkg/clients/user/grpc"
)

type UserService struct {
	user grpc.UserProviderer
}

func NewUserService(user grpc.UserProviderer) *UserService {
	return &UserService{user: user}
}

func (u *UserService) Create(in CreateIn) CreateOut {
	err := u.user.Create(in.User)
	if err != nil {
		return CreateOut{"Не удалось создать пользователя", err}
	}
	return CreateOut{
		Message: "Пользователь успешно создан",
		Error:   nil,
	}
}

func (u *UserService) GetByUsername(in GetIn) GetOut {
	user, err := u.user.Profile(in.Username)
	if err != nil {
		return GetOut{models.User{}, err}
	}
	return GetOut{
		User:  user,
		Error: nil,
	}
}

func (u *UserService) List() ListOut {
	users, err := u.user.List()
	if err != nil {
		return ListOut{nil, err}
	}
	return ListOut{users, err}
}
