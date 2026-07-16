package domain

import (
	"context"
	"database/sql"
	"industrial-supply-store/internal/model/entity"
)

type ProductRepository interface {
	FindAll(ctx context.Context) ([]entity.Product, error)
	Create(ctx context.Context, product *entity.Product) error
	Update(ctx context.Context, product *entity.Product) error
	Delete(ctx context.Context, id int) error

	GetAllProducts(ctx context.Context) ([]entity.ProductWithSupplier, error)
	GetProductByID(ctx context.Context, id int) (entity.ProductWithSupplier, error)
	CreateProduct(ctx context.Context, product entity.ProductWithSupplier) error
	UpdateProduct(ctx context.Context, product entity.ProductWithSupplier) error
	DeleteProduct(ctx context.Context, id int) error
	UpdateStock(ctx context.Context, tx *sql.Tx, productID int, qty int) error
}
