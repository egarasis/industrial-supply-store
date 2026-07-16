package domain

import (
	"context"
	"database/sql"
	"industrial-supply-store/internal/model/entity"
)

type OrderRepository interface {

	// Checkout
	CreateOrder(ctx context.Context, tx *sql.Tx, order entity.Order) (int, error)
	CreateOrderItem(ctx context.Context, tx *sql.Tx, item entity.OrderItem) error
	UpdateOrderTotal(ctx context.Context, tx *sql.Tx, orderID int, total float64) error
	Checkout(ctx context.Context, userID int, cart []entity.CartItem) error

	// Customer
	GetOrdersByUserID(ctx context.Context, userID int) ([]entity.Order, error)
	GetOrderByID(ctx context.Context, orderID int) (entity.Order, error)
	GetOrderItems(ctx context.Context, orderID, userID int) ([]entity.OrderItem, error)

	// Admin
	GetAllOrders(ctx context.Context) ([]entity.OrderWithUser, error)
	UpdateOrderStatus(ctx context.Context, orderID int, status string) error
	GetOrdersByStatus(ctx context.Context, status string) ([]entity.OrderWithUser, error)
}
