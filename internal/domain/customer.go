package domain

import (
	"context"
	"industrial-supply-store/internal/model/entity"
)

type CustomerHandler interface {
	Run(int)
}

<<<<<<< HEAD
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
=======
type OrderUsecase interface {
	Checkout(ctx context.Context, userID int, cart []entity.CartItem) error
	GetMyOrders(ctx context.Context, userID int) ([]entity.Order, error)
	GetOrderDetail(ctx context.Context, orderID, userID int) ([]entity.OrderItem, error)
	UpdateOrderStatus(ctx context.Context, orderID int, status string) error
	GetAllProducts(
		ctx context.Context,
	) ([]entity.ProductWithSupplier, error)
}
>>>>>>> a9d2308fdc4245458fd69dd2b7b286b0217a42fc
