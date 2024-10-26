package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Product struct {
	ID        int       `json:"id" db:"id"`
	Gender    string    `json:"gender" db:"gender"`
	Style     string    `json:"style" db:"style"`
	Size      string    `json:"size" db:"size"`
	Price     float64   `json:"price" db:"price"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Order struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Address   string    `json:"address" db:"address"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type OrderItem struct {
	ID        int       `json:"id" db:"id"`
	OrderID   int       `json:"order_id" db:"order_id"`
	ProductID int       `json:"product_id" db:"product_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreateOrderRequest struct {
	UserID     int         `json:"user_id"`
	Address    string      `json:"address"`
	OrderItems []OrderItem `json:"order_items"`
}

type ListOrderRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Status    string `json:"status"`
	PerPage   int    `json:"per_page"`
	Page      int    `json:"page"`
}
