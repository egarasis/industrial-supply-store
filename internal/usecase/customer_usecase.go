package usecase

import (
	"context"
	"database/sql"
	"errors"

	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
)

type orderUsecase struct {
	db          *sql.DB
	orderRepo   domain.OrderRepository
	productRepo domain.ProductRepository
	userRepo    domain.UserRepository // Tambahkan ini
}

func NewCustomerUsecase(
	db *sql.DB,
	orderRepo domain.OrderRepository,
	productRepo domain.ProductRepository,
	userRepo domain.UserRepository, // Tambahkan parameter ini
) domain.OrderUsecase {
	return &orderUsecase{
		db:          db,
		orderRepo:   orderRepo,
		productRepo: productRepo,
		userRepo:    userRepo, // Tambahkan ini
	}
}

func (u *orderUsecase) UpdateProfile(ctx context.Context, userID int, fullName string, companyName string) error {
	profile := entity.UserProfile{
		UserID:      userID,
		FullName:    fullName,
		CompanyName: companyName,
	}
	return u.userRepo.UpdateProfile(ctx, profile)
}

func (u *orderUsecase) Checkout(
	ctx context.Context,
	userID int,
	cart []entity.CartItem,
) error {

	if len(cart) == 0 {
		return errors.New("cart is empty")
	}

	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	order := entity.Order{
		UserID:     userID,
		TotalPrice: 0,
		Status:     "Pending",
	}

	orderID, err := u.orderRepo.CreateOrder(ctx, tx, order)
	if err != nil {
		return err
	}

	var total float64

	for _, cartItem := range cart {

		product, err := u.productRepo.GetProductByID(ctx, cartItem.ProductID)
		if err != nil {
			return err
		}

		if product.Stock < cartItem.Quantity {
			return errors.New("stock is not enough")
		}

		subtotal := product.Price * float64(cartItem.Quantity)

		orderItem := entity.OrderItem{
			OrderID:   orderID,
			ProductID: product.ID,
			Quantity:  cartItem.Quantity,
			Subtotal:  subtotal,
		}

		err = u.orderRepo.CreateOrderItem(ctx, tx, orderItem)
		if err != nil {
			return err
		}

		err = u.productRepo.UpdateStock(
			ctx,
			tx,
			product.ID,
			cartItem.Quantity,
		)

		if err != nil {
			return err
		}

		total += subtotal
	}

	err = u.orderRepo.UpdateOrderTotal(
		ctx,
		tx,
		orderID,
		total,
	)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (u *orderUsecase) GetMyOrders(
	ctx context.Context,
	userID int,
) ([]entity.Order, error) {

	return u.orderRepo.GetOrdersByUserID(ctx, userID)
}

func (u *orderUsecase) GetOrderDetail(
	ctx context.Context,
	orderID int,
) ([]entity.OrderItem, error) {

	return u.orderRepo.GetOrderItems(ctx, orderID)
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
