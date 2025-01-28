package pg

import (
	"context"
	"fmt"
	"webproj/internal/domain"
)

type ItemPgRepo struct {
	db *DB
}

func NewItemPgRepo(db *DB) *ItemPgRepo {
	return &ItemPgRepo{db: db}
}
func (r *ItemPgRepo) Create(ctx context.Context, item *domain.Item) error {
	query := `
        INSERT INTO items (name,description,price,stock)
        VALUES ($1,$2,$3,$4)
    `
	_, err := r.db.Exec(ctx, query,
		item.Name,
		item.Description,
		item.Price,
		item.Stock,
	)
	if err != nil {
		return fmt.Errorf("Failed to create item")
	}
	return nil
}
func (r *ItemPgRepo) GetById(ctx context.Context, id int) (*domain.Item, error) {
	var itemDB domain.Item
	query := `SELECT id,name,description,price,stock FROM items WHERE id=$1`
	err := r.db.QueryRow(ctx, query, id).Scan(&itemDB.Id, &itemDB.Name, &itemDB.Description, &itemDB.Price, &itemDB.Stock)
	if err != nil {
		return nil, fmt.Errorf("failed to get item by id")
	}
	return &itemDB, nil
}

func (r *ItemPgRepo) GetAll(ctx context.Context) ([]domain.Item, error) {
	items := make([]domain.Item, 0, 5)
	query := `SELECT * FROM items`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to select items")
	}
	defer rows.Close()
	for rows.Next() {
		var item domain.Item
		err := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			&item.Price,
			&item.Stock,
		)
		if err != nil {
			return nil, fmt.Errorf("failed scan from row")
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *ItemPgRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM items WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
