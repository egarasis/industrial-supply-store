package domain

import (
	"context"
	"database/sql"
	"industrial-supply-store/internal/model/entity"
)

// OrderRepository adalah interface untuk akses data order
type OrderRepository interface {
	CreateOrder(ctx context.Context, tx *sql.Tx, order entity.Order) (int, error)
	CreateOrderItem(ctx context.Context, tx *sql.Tx, item entity.OrderItem) error
	UpdateOrderTotal(ctx context.Context, tx *sql.Tx, orderID int, total float64) error
	GetOrdersByUserID(ctx context.Context, userID int) ([]entity.Order, error)
	GetOrderByID(ctx context.Context, orderID int) (entity.Order, error)
	GetOrderItems(ctx context.Context, orderID int) ([]entity.OrderItem, error)
	UpdateOrderStatus(ctx context.Context, orderID int, status string) error // Tambahkan ini kalau dipakai
}

// OrderUsecase adalah interface untuk logika bisnis
type OrderUsecase interface {
	Checkout(ctx context.Context, userID int, cart []entity.CartItem) error
	GetMyOrders(ctx context.Context, userID int) ([]entity.Order, error)
	GetOrderDetail(ctx context.Context, orderID int) ([]entity.OrderItem, error)
	UpdateOrderStatus(ctx context.Context, orderID int, status string) error
	GetAllProducts(ctx context.Context) ([]entity.ProductWithSupplier, error)
	UpdateProfile(ctx context.Context, userID int, fullName string, companyName string) error
}
