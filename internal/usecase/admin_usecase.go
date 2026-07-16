package usecase

import (
	"context"

	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
)

type adminUsecase struct {
	productRepo domain.ProductRepository
}

func NewAdminUsecase(productRepo domain.ProductRepository) domain.AdminUsecase {
	return &adminUsecase{
		productRepo: productRepo,
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

func (u *adminUsecase) AssignCategoryToProduct(ctx context.Context, productID, categoryID int) error {
	return u.productRepo.AssignCategory(ctx, productID, categoryID)
}

func (u *adminUsecase) GetUserReport(ctx context.Context) ([]entity.UserReport, error) {
	return u.productRepo.GetUserReport(ctx)
}

func (u *adminUsecase) GetStockReport(ctx context.Context) ([]entity.StockReport, error) {
	return u.productRepo.GetStockReport(ctx)
}
