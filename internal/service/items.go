package service

import (
	"context"
	"webproj/internal/domain"
	"webproj/internal/repository"
)

type ItemService struct {
	repo repository.Items
}

func NewItemService(repo repository.Items) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

func (s *ItemService) NewItem(ctx context.Context, itemValues *ItemValues) error {
	for _, v := range itemValues.Items {
		item := domain.Item{
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			Stock:       v.Stock,
		}
		err := s.repo.Create(ctx, &item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ItemService) GetAll(ctx context.Context) ([]domain.Item, error) {
	return s.repo.GetAll(ctx)
}
func (s *ItemService) Get(ctx context.Context, itemId int) (*domain.Item, error) {
	item, err := s.repo.GetById(ctx, itemId)
	if err != nil {
		return nil, err
	}
	return item, nil
}
func (s *ItemService) Delete(ctx context.Context, itemId int) error {
	return s.repo.Delete(ctx, itemId)
}
