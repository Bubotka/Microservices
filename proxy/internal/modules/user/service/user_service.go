package service

import (
	"context"
	"geo/internal/db/adapter"
	"main/internal/modules/user/storage"
)

type User struct {
	storage storage.UserRepository
}

func NewUserService(storage storage.UserRepository) *User {
	return &User{storage: storage}
}

func (u *User) Create(in CreateIn) CreateOut {
	err := u.storage.Create(context.Background(), in.User)
	if err != nil {
		return CreateOut{"", err}
	}
	return CreateOut{"Пользователь успешно создан", nil}
}

func (u *User) GetByID(in GetIn) GetOut {
	user, err := u.storage.GetByID(context.Background(), in.Id)
	return GetOut{
		User:  user,
		Error: err,
	}
}

func (u *User) Update(in UpdateIn) UpdateOut {
	err := u.storage.Update(context.Background(), in.User)
	if err != nil {
		return UpdateOut{"", err}
	}
	return UpdateOut{"Пользователь успешно обновлён", nil}
}

func (u *User) Delete(in DeleteIn) DeleteOut {
	err := u.storage.Delete(context.Background(), in.Id)
	if err != nil {
		return DeleteOut{"", err}
	}
	return DeleteOut{"Пользователь успешно удалён", nil}
}

func (u *User) List(in ListIn) ListOut {
	user, err := u.storage.List(context.Background(), adapter.Condition{
		LimitOffset: &adapter.LimitOffset{
			Offset: int64(in.Offset),
			Limit:  int64(in.Limit),
		},
	})
	return ListOut{
		User:  user,
		Error: err,
	}
}

func (u *User) Count() CountOut {
	count, err := u.storage.Count(context.Background(), adapter.Condition{})
	return CountOut{
		Count: count,
		Error: err,
	}
}
