package domain

import (
	"context"
	"database/sql"
	"industrial-supply-store/internal/model/entity"
)

type CustomerHandler interface {
	Run(int)
}

// OrderCustomerRepository tetap di sini karena spesifik untuk domain customer
type OrderCustomerRepository interface {
	// Checkout
	CreateOrder(ctx context.Context, tx *sql.Tx, order entity.Order) (int, error)
	CreateOrderItem(ctx context.Context, tx *sql.Tx, item entity.OrderItem) error
	UpdateOrderTotal(ctx context.Context, tx *sql.Tx, orderID int, total float64) error

	// Customer
	GetOrdersByUserID(ctx context.Context, userID int) ([]entity.Order, error)
	GetOrderByID(ctx context.Context, orderID int) (entity.Order, error)
	GetOrderItems(ctx context.Context, orderID int) ([]entity.OrderItem, error)
}

// OrderUsecase sudah didefinisikan di internal/domain/order.go,
// jadi tidak perlu didefinisikan lagi di sini agar tidak error redeclared.
