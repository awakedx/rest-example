package service

import (
	"context"
	"fmt"
	"time"
	"webproj/internal/domain"
	"webproj/internal/repository"

	"github.com/google/uuid"
)

type OrderService struct {
	repoOrder repository.Orders
	repoItems repository.Items
}

func NewOrderService(repoOrder repository.Orders, repoItems repository.Items) *OrderService {
	return &OrderService{
		repoOrder: repoOrder,
		repoItems: repoItems,
	}
}

func (s *OrderService) MakeOrder(ctx context.Context, input *InputOrder) (int, error) {
	var totalSum float64 = 0
	ItemPrices := make(map[int]float64)
	for _, v := range input.Items {
		i, err := s.repoItems.GetById(ctx, v.ItemId)
		if err != nil {
			return 0, err
		}
		if i.Stock < v.Quantity {
			return 0, fmt.Errorf("out of stock %s (requsted: %d, available: %d", i.Name, v.Quantity, i.Stock)
		}
		totalSum += i.Price * float64(v.Quantity)
		ItemPrices[i.Id] = i.Price
	}

	order := domain.Order{
		UserId:     input.UserId,
		OrderDate:  time.Now(),
		TotalPrice: totalSum,
		Items:      input.Items,
	}
	OrderId, err := s.repoOrder.Create(ctx, &order, ItemPrices)
	if err != nil {
		return 0, err
	}
	return OrderId, err
}
func (s *OrderService) GetById(ctx context.Context, orderId int, userId uuid.UUID) (*domain.Order, error) {
	order, err := s.repoOrder.GetById(ctx, orderId)
	if err != nil || order.UserId != userId {
		return nil, fmt.Errorf("invalid order id")
	}
	return order, nil

}
func (s *OrderService) GetAll(ctx context.Context, userId uuid.UUID) ([]domain.Order, error) {
	orders, err := s.repoOrder.GetAllByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
