package domain

import (
	"context"
	"industrial-supply-store/internal/model/entity"
)

type CustomerHandler interface {
	Run(int)
}

type OrderUsecase interface {
	Checkout(ctx context.Context, userID int, cart []entity.CartItem) error
	GetMyOrders(ctx context.Context, userID int) ([]entity.Order, error)
	GetOrderDetail(ctx context.Context, orderID, userID int) ([]entity.OrderItem, error)
	UpdateOrderStatus(ctx context.Context, orderID int, status string) error
	GetAllProducts(
		ctx context.Context,
	) ([]entity.ProductWithSupplier, error)
}
