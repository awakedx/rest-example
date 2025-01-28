package service

import (
	"context"
	"time"
	"webproj/internal/domain"
	"webproj/internal/repository"

	"github.com/google/uuid"
)

type Services struct {
	Users  Users
	Items  Items
	Orders Orders
}

type Deps struct {
	Repos          *repository.Repositories
	AccessTokenTTL time.Duration
}

func NewServices(deps Deps) *Services {
	return &Services{
		Users:  NewUsersService(deps.Repos.Users, deps.AccessTokenTTL),
		Items:  NewItemService(deps.Repos.Items),
		Orders: NewOrderService(deps.Repos.Orders, deps.Repos.Items),
	}
}

type SignUpInput struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}

type SignInInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type InputItem struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"desc" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
}

type ItemValues struct {
	Items []InputItem `json:"items" validate:"required"`
}

type InputOrder struct {
	UserId uuid.UUID          `json:"userId"`
	Items  []domain.OrderItem `json:"items"`
}

type Users interface {
	SignUp(ctx context.Context, input *SignUpInput) error
	SignIn(ctx context.Context, input *SignInInput) (string, error)
	DeleteUser(ctx context.Context, input uuid.UUID) error
	GetById(ctx context.Context, input uuid.UUID) (*domain.User, error)
	GetAccTokenTTL() time.Duration
}

type Items interface {
	NewItem(ctx context.Context, input *ItemValues) error
	GetAll(ctx context.Context) ([]domain.Item, error)
	Get(ctx context.Context, itemId int) (*domain.Item, error)
	Delete(ctx context.Context, itemId int) error
}

type Orders interface {
	MakeOrder(ctx context.Context, input *InputOrder) (int, error)
	GetById(ctx context.Context, orderId int, userId uuid.UUID) (*domain.Order, error)
	GetAll(ctx context.Context, input uuid.UUID) ([]domain.Order, error)
}
