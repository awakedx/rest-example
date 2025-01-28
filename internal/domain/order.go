package domain

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	Id         int         `json:"id"`
	UserId     uuid.UUID   `json:"userId"`
	OrderDate  time.Time   `json:"orderDate"`
	TotalPrice float64     `json:"totalPrice"`
	Items      []OrderItem `json:"items"`
}

type OrderItem struct {
	ItemId   int     `json:"itemId"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
