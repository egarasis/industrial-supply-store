package domain

import (
	"context"
	"industrial-supply-store/internal/model/entity"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category entity.Category) error
	GetAllCategories(ctx context.Context) ([]entity.Category, error)
	AssignCategory(ctx context.Context, pc entity.ProductCategory) error
}
