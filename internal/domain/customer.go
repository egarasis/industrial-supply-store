package domain

import (
	"context"
	"database/sql"
	"industrial-supply-store/internal/model/entity"
)

type CustomerHandler interface {
	Run(int)
}

// type AdminUsecase interface {
// 	ListProducts(ctx context.Context) ([]entity.Product, error)
// 	AddProduct(ctx context.Context, product *entity.Product) error
// 	UpdateProduct(ctx context.Context, product *entity.Product) error
// 	DeleteProduct(ctx context.Context, id int) error
// }

// type ProductRepository interface {
// 	FindAll(ctx context.Context) ([]entity.Product, error)
// 	Create(ctx context.Context, product *entity.Product) error
// 	Update(ctx context.Context, product *entity.Product) error
// 	Delete(ctx context.Context, id int) error

// 	GetAllProducts(ctx context.Context) ([]entity.ProductWithSupplier, error)
// 	GetProductByID(ctx context.Context, id int) (entity.ProductWithSupplier, error)
// 	CreateProduct(ctx context.Context, product entity.ProductWithSupplier) error
// 	UpdateProduct(ctx context.Context, product entity.ProductWithSupplier) error
// 	DeleteProduct(ctx context.Context, id int) error
// 	UpdateStock(ctx context.Context, tx *sql.Tx, productID int, qty int) error
// }

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

type OrderUsecase interface {
	Checkout(ctx context.Context, userID int, cart []entity.CartItem) error
	GetMyOrders(ctx context.Context, userID int) ([]entity.Order, error)
	GetOrderDetail(ctx context.Context, orderID int) ([]entity.OrderItem, error)
	UpdateOrderStatus(ctx context.Context, orderID int, status string) error
	GetAllProducts(
		ctx context.Context,
	) ([]entity.ProductWithSupplier, error)
}
