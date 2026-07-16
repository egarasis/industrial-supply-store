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
<<<<<<< HEAD
	AssignCategoryToProduct(ctx context.Context, productID, categoryID int) error
	GetUserReport(ctx context.Context) ([]entity.UserReport, error)
	GetStockReport(ctx context.Context) ([]entity.StockReport, error)
}
=======
>>>>>>> a9d2308fdc4245458fd69dd2b7b286b0217a42fc

	ListOrders(ctx context.Context) ([]entity.OrderWithUser, error)
	GetOrdersByStatus(ctx context.Context, status string) ([]entity.OrderWithUser, error)
	UpdateOrderStatus(ctx context.Context, orderID int, status string) error
}