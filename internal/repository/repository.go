package repository

import (
	"context"
	"webproj/internal/domain"
	pg "webproj/internal/repository/PG"

	"github.com/google/uuid"
)

type Users interface {
	Create(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	Get(ctx context.Context, emailOrId interface{}) (*domain.User, error)
}

type Items interface {
	Create(ctx context.Context, item *domain.Item) error
	GetAll(ctx context.Context) ([]domain.Item, error)
	GetById(ctx context.Context, id int) (*domain.Item, error)
	Delete(ctx context.Context, id int) error
}

type Orders interface {
	Create(ctx context.Context, order *domain.Order, itemPrices map[int]float64) (int, error)
	GetAllByUserId(ctx context.Context, userId uuid.UUID) ([]domain.Order, error)
	GetById(ctx context.Context, orderId int) (*domain.Order, error)
}

type Repositories struct {
	Users  Users
	Items  Items
	Orders Orders
}

func NewRepositories(db *pg.DB) *Repositories {
	return &Repositories{
		Users:  pg.NewUserPgRepo(db),
		Items:  pg.NewItemPgRepo(db),
		Orders: pg.NewOrderPgRepo(db),
	}
}
