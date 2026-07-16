package usecase

import (
	"context"

	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
)

type adminUsecase struct {
	productRepo domain.ProductRepository
	orderRepo   domain.OrderRepository
}

func NewAdminUsecase(productRepo domain.ProductRepository, orderRepo domain.OrderRepository) domain.AdminUsecase {
	return &adminUsecase{
		productRepo: productRepo,
		orderRepo:   orderRepo,
	}
}

func (u *adminUsecase) ListProducts(ctx context.Context) ([]entity.Product, error) {
	return u.productRepo.FindAll(ctx)
}

func (u *adminUsecase) AddProduct(ctx context.Context, product *entity.Product) error {
	return u.productRepo.Create(ctx, product)
}

func (u *adminUsecase) UpdateProduct(ctx context.Context, product *entity.Product) error {
	return u.productRepo.Update(ctx, product)
}

func (u *adminUsecase) DeleteProduct(ctx context.Context, id int) error {
	return u.productRepo.Delete(ctx, id)
}

<<<<<<< HEAD
func (u *adminUsecase) AssignCategoryToProduct(ctx context.Context, productID, categoryID int) error {
	return u.productRepo.AssignCategory(ctx, productID, categoryID)
}

func (u *adminUsecase) GetUserReport(ctx context.Context) ([]entity.UserReport, error) {
	return u.productRepo.GetUserReport(ctx)
}

func (u *adminUsecase) GetStockReport(ctx context.Context) ([]entity.StockReport, error) {
	return u.productRepo.GetStockReport(ctx)
}
=======
func (u *adminUsecase) ListOrders(ctx context.Context) ([]entity.OrderWithUser, error) {
	return u.orderRepo.GetAllOrders(ctx)
}

func (u *adminUsecase) GetOrdersByStatus(ctx context.Context, status string) ([]entity.OrderWithUser, error) {
	return u.orderRepo.GetOrdersByStatus(ctx, status)
}

func (u *adminUsecase) UpdateOrderStatus(ctx context.Context, orderID int, status string) error {
	return u.orderRepo.UpdateOrderStatus(ctx, orderID, status)
}
>>>>>>> a9d2308fdc4245458fd69dd2b7b286b0217a42fc
