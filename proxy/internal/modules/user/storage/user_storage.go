package storage

import (
	"context"
	"fmt"
	"github.com/Bubotka/Microservices/geo/pkg/db/adapter"
	"github.com/Bubotka/Microservices/proxy/internal/models"
	"strconv"
)

type UserStorage struct {
	adapter adapter.SqlAdapterer
}

func NewUserStorage(adapter adapter.SqlAdapterer) *UserStorage {
	return &UserStorage{adapter: adapter}
}

func (u *UserStorage) Create(ctx context.Context, user models.User) error {
	err := u.adapter.Create(ctx, user)

	return err
}

func (u *UserStorage) GetByID(ctx context.Context, id string) (models.User, error) {
	var user []models.User
	intId, _ := strconv.Atoi(id)
	err := u.adapter.List(ctx, &user, models.User{}, adapter.Condition{
		Equal: map[string]interface{}{
			"id": intId,
		},
	})

	fmt.Println(user)

	return user[0], err
}

func (u *UserStorage) Update(ctx context.Context, user models.User) error {
	err := u.adapter.Update(ctx, user, adapter.Condition{
		Equal: map[string]interface{}{
			"id": user.ID,
		}}, map[string]interface{}{"id": user.ID}, map[string]interface{}{"username": user.Username}, map[string]interface{}{"password": user.Password})
	return err
}

func (u *UserStorage) Delete(ctx context.Context, id int) error {
	err := u.adapter.Update(ctx, models.User{}, adapter.Condition{
		Equal: map[string]interface{}{
			"id": id,
		}}, map[string]interface{}{"is_delete": true})
	return err
}

func (u *UserStorage) List(ctx context.Context, c adapter.Condition) ([]models.User, error) {
	var users []models.User
	err := u.adapter.List(ctx, &users, models.User{}, c)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserStorage) Count(ctx context.Context, c adapter.Condition) (int, error) {
	count, err := u.adapter.GetCount(ctx, models.User{}, c)
	if err != nil {
		return 0, err
	}
	return count, nil
}
