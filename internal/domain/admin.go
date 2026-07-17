package domain

import (
	"context"
	"industrial-supply-store/internal/model/entity"
)

type AdminHandler interface {
	Run()
}

type AdminUsecase interface {
	ListProducts(ctx context.Context) ([]entity.Product, error)
	AddProduct(ctx context.Context, product *entity.Product) error
	UpdateProduct(ctx context.Context, product *entity.Product) error
	DeleteProduct(ctx context.Context, id int) error

	ListOrders(ctx context.Context) ([]entity.OrderWithUser, error)
	GetOrdersByStatus(ctx context.Context, status string) ([]entity.OrderWithUser, error)
	UpdateOrderStatus(ctx context.Context, orderID int, status string) error

	NewListProducts(ctx context.Context) ([]entity.ProductWithSupplier, error)
	NewCreateProduct(ctx context.Context, product entity.ProductWithSupplier) error
	NewUpdateProduct(ctx context.Context, product entity.ProductWithSupplier) error
	NewDeleteProduct(ctx context.Context, id int) error

	CreateCategory(ctx context.Context, category entity.Category) error
	GetAllCategories(ctx context.Context) ([]entity.Category, error)
	AssignCategory(ctx context.Context, pc entity.ProductCategory) error

	GetTopUser(ctx context.Context) ([]entity.UserReport, error)
	GetOutOfStock(ctx context.Context) ([]entity.StockReport, error)
}
