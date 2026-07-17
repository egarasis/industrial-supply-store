package usecase

import (
	"context"
	"errors"

	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
)

type adminUsecase struct {
	productRepo  domain.ProductRepository
	orderRepo    domain.OrderRepository
	categoryRepo domain.CategoryRepository
	userRepo     domain.UserRepository
}

func NewAdminUsecase(
	productRepo domain.ProductRepository,
	orderRepo domain.OrderRepository,
	categoryRepo domain.CategoryRepository,
	userRepo domain.UserRepository,
) domain.AdminUsecase {
	return &adminUsecase{
		productRepo:  productRepo,
		orderRepo:    orderRepo,
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
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

func (u *adminUsecase) ListOrders(ctx context.Context) ([]entity.OrderWithUser, error) {
	return u.orderRepo.GetAllOrders(ctx)
}

func (u *adminUsecase) GetOrdersByStatus(ctx context.Context, status string) ([]entity.OrderWithUser, error) {
	return u.orderRepo.GetOrdersByStatus(ctx, status)
}

func (u *adminUsecase) UpdateOrderStatus(ctx context.Context, orderID int, status string) error {
	return u.orderRepo.UpdateOrderStatus(ctx, orderID, status)
}

func (u *adminUsecase) NewCreateProduct(ctx context.Context, product entity.ProductWithSupplier) error {

	if product.SupplierID <= 0 {
		return errors.New("supplier is required")
	}

	if product.ProductName == "" {
		return errors.New("product name is required")
	}

	if product.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	if product.Stock < 0 {
		return errors.New("stock cannot be negative")
	}

	return u.productRepo.CreateProduct(ctx, product)
}

func (u *adminUsecase) NewUpdateProduct(ctx context.Context, product entity.ProductWithSupplier) error {

	if product.ID <= 0 {
		return errors.New("invalid product id")
	}

	if product.SupplierID <= 0 {
		return errors.New("supplier is required")
	}

	if product.ProductName == "" {
		return errors.New("product name is required")
	}

	if product.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	if product.Stock < 0 {
		return errors.New("stock cannot be negative")
	}

	return u.productRepo.UpdateProduct(ctx, product)
}

func (u *adminUsecase) NewDeleteProduct(ctx context.Context, id int) error {

	if id <= 0 {
		return errors.New("invalid product id")
	}

	return u.productRepo.DeleteProduct(ctx, id)
}

func (u *adminUsecase) NewListProducts(ctx context.Context) ([]entity.ProductWithSupplier, error) {

	products, err := u.productRepo.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (u *adminUsecase) CreateCategory(ctx context.Context, category entity.Category) error {

	if category.CategoryName == "" {
		return errors.New("category name is required")
	}

	return u.categoryRepo.CreateCategory(ctx, category)
}

func (u *adminUsecase) GetAllCategories(ctx context.Context) ([]entity.Category, error) {

	return u.categoryRepo.GetAllCategories(ctx)
}

func (u *adminUsecase) AssignCategory(ctx context.Context, pc entity.ProductCategory) error {

	if pc.ProductID <= 0 {
		return errors.New("invalid product")
	}

	if pc.CategoryID <= 0 {
		return errors.New("invalid category")
	}

	return u.categoryRepo.AssignCategory(ctx, pc)
}

func (u *adminUsecase) GetTopUser(ctx context.Context) ([]entity.UserReport, error) {
	return u.userRepo.GetTopUser(ctx)
}

func (u *adminUsecase) GetOutOfStock(ctx context.Context) ([]entity.StockReport, error) {
	return u.productRepo.GetOutOfStock(ctx)
}
