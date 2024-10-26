package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Rawipass/product-service/config"
	"github.com/Rawipass/product-service/models"
	"github.com/jackc/pgx/v4"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository() *ProductRepository {
	repo := ProductRepository{}
	return &repo
}

func (p *ProductRepository) ListProducts(gender, style, size string, page, perPage int) ([]models.Product, error) {
	var products []models.Product

	query := "SELECT id, gender, style, size, price, created_at FROM products WHERE 1=1"
	args := []interface{}{}
	argIndex := 1

	if gender != "" {
		query += fmt.Sprintf(" AND gender = $%d", argIndex)
		args = append(args, gender)
		argIndex++
	}

	if style != "" {
		query += fmt.Sprintf(" AND style = $%d", argIndex)
		args = append(args, style)
		argIndex++
	}

	if size != "" {
		query += fmt.Sprintf(" AND size = $%d", argIndex)
		args = append(args, size)
		argIndex++
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, perPage, (page-1)*perPage)
	rows, err := config.DB.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Gender, &product.Style, &product.Size, &product.Price, &product.CreatedAt)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepository) CreateOrder(orderRequest models.CreateOrderRequest) (int, error) {
	var orderID int
	tx, err := config.DB.Begin(context.Background())
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(context.Background())

	orderQuery := `INSERT INTO orders (user_id, address, status, created_at) 
                   VALUES ($1, $2, 'placed_order', NOW()) RETURNING id`
	err = tx.QueryRow(context.Background(), orderQuery, orderRequest.UserID, orderRequest.Address).Scan(&orderID)
	if err != nil {
		return 0, err
	}

	err = p.createOrderItems(tx, orderID, orderRequest.OrderItems)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return 0, err
	}

	return orderID, nil
}

func (p *ProductRepository) createOrderItems(tx pgx.Tx, orderID int, orderItems []models.OrderItem) error {
	for _, item := range orderItems {
		itemQuery := `INSERT INTO order_items (order_id, product_id, quantity, created_at) 
                       VALUES ($1, $2, $3, NOW())`

		_, err := tx.Exec(context.Background(), itemQuery, orderID, item.ProductID, item.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *ProductRepository) ListOrders(start_date, end_date, status string, page, perPage int) ([]models.Order, error) {
	var orders []models.Order

	query := `SELECT id, user_id, address, status, created_at 
              FROM orders 
              WHERE 1=1`

	args := []interface{}{}
	argIndex := 1

	if start_date != "" && end_date != "" {
		query += fmt.Sprintf(" AND created_at BETWEEN TO_TIMESTAMP($%d, 'YYYY-MM-DD HH24:MI:SS') AND TO_TIMESTAMP($%d, 'YYYY-MM-DD HH24:MI:SS')", argIndex, argIndex+1)
		args = append(args, start_date, end_date)
		argIndex += 2
	}

	if status != "" {
		query += fmt.Sprintf(" AND status = $%d", argIndex)
		args = append(args, status)
		argIndex++
	}
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, perPage, (page-1)*perPage)

	rows, err := config.DB.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.Address, &order.Status, &order.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
