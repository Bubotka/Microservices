package storage

import (
	"context"
	"geo/internal/db/adapter"
	"main/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.36.0 --name=UserRepository
type UserRepository interface {
	Create(ctx context.Context, user models.User) error
	GetByID(ctx context.Context, id string) (models.User, error)
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, c adapter.Condition) ([]models.User, error)
	Count(ctx context.Context, c adapter.Condition) (int, error)
}
