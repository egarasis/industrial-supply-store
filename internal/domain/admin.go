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
}

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
