package usecase

import (
	"context"
	"errors"

	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
)

type orderUsecase struct {
	orderRepo   domain.OrderRepository
	productRepo domain.ProductRepository
}

func NewCustomerUsecase(
	orderRepo domain.OrderRepository,
	productRepo domain.ProductRepository,
) domain.OrderUsecase {
	return &orderUsecase{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (u *orderUsecase) Checkout(
	ctx context.Context,
	userID int,
	cart []entity.CartItem,
) error {

	// Validation
	if len(cart) == 0 {
		return errors.New("cart is empty")
	}

	return u.orderRepo.Checkout(
		ctx,
		userID,
		cart,
	)
}

func (u *orderUsecase) GetMyOrders(
	ctx context.Context,
	userID int,
) ([]entity.Order, error) {

	return u.orderRepo.GetOrdersByUserID(ctx, userID)
}

func (u *orderUsecase) GetOrderDetail(
	ctx context.Context,
	orderID, userID int,
) ([]entity.OrderItem, error) {

	return u.orderRepo.GetOrderItems(ctx, orderID, userID)
}

func (u *orderUsecase) UpdateOrderStatus(
	ctx context.Context,
	orderID int,
	status string,
) error {

	return u.orderRepo.UpdateOrderStatus(ctx, orderID, status)
}

func (u *orderUsecase) GetAllProducts(
	ctx context.Context,
) ([]entity.ProductWithSupplier, error) {

	return u.productRepo.GetAllProducts(ctx)
}
