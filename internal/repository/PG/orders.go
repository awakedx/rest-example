package pg

import (
	"context"
	"fmt"
	"webproj/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type OrderPgRepo struct {
	db *DB
}

func NewOrderPgRepo(db *DB) *OrderPgRepo {
	return &OrderPgRepo{
		db: db,
	}
}
func (r *OrderPgRepo) Create(ctx context.Context, order *domain.Order, itemPrices map[int]float64) (int, error) {
	var idOrder int
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	query := `
    INSERT INTO orders (user_id,order_date,total_price)
    VALUES($1,$2,$3)
    RETURNING id
    `
	err = tx.QueryRow(ctx, query,
		order.UserId,
		order.OrderDate,
		order.TotalPrice,
	).Scan(&idOrder)
	if err != nil {
		tx.Rollback(ctx)
		return 0, err
	}
	query = `
    INSERT INTO order_items (order_id,item_id,quantity,price)
    VALUES($1,$2,$3,$4)
    `
	for _, v := range order.Items {
		_, err = tx.Exec(ctx, query,
			idOrder,
			v.ItemId,
			v.Quantity,
			itemPrices[v.ItemId],
		)
		if err != nil {
			tx.Rollback(ctx)
			return 0, err
		}
		query2 := `UPDATE items SET stock=stock-$1 WHERE id=$2`
		_, err := tx.Exec(ctx, query2, v.Quantity, v.ItemId)
		if err != nil {
			tx.Rollback(ctx)
			return 0, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}
	return idOrder, nil
}
func (r *OrderPgRepo) GetById(ctx context.Context, orderId int) (*domain.Order, error) {
	query := "SELECT * FROM orders WHERE id=$1"
	var order domain.Order
	err := r.db.QueryRow(ctx, query, orderId).Scan(&order.Id, &order.UserId, &order.OrderDate, &order.TotalPrice)
	if err != nil {
		return nil, fmt.Errorf("failed to scan to order")
	}
	order.Items, err = r.GetOrderItemsById(ctx, orderId)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderPgRepo) GetOrderItemsById(ctx context.Context, orderId int) ([]domain.OrderItem, error) {
	queryItems := `SELECT item_id,quantity,price FROM order_items WHERE order_id=$1`
	rowsItems, err := r.db.Query(ctx, queryItems, orderId)
	if err != nil {
		return nil, err
	}
	items := make([]domain.OrderItem, 0, 2)
	for rowsItems.Next() {
		var orderItem domain.OrderItem
		err := rowsItems.Scan(
			&orderItem.ItemId,
			&orderItem.Quantity,
			&orderItem.Price,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, orderItem)
	}
	return items, nil
}

func (r *OrderPgRepo) GetAllByUserId(ctx context.Context, userId uuid.UUID) ([]domain.Order, error) {
	query := `SELECT id,order_date,total_price FROM orders WHERE user_id=$1`

	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	orders := make([]domain.Order, 0, 3)
	for rows.Next() {
		order := domain.Order{UserId: userId, Items: []domain.OrderItem{}}
		err := rows.Scan(
			&order.Id,
			&order.OrderDate,
			&order.TotalPrice,
		)
		if err != nil {
			return nil, err
		}
		order.Items, err = r.GetOrderItemsById(ctx, order.Id)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
