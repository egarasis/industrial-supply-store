package db

import (
	"context"
	"database/sql"
	"errors"
	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
)

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(
	db *sql.DB,
) domain.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) GetAllCategories(ctx context.Context) ([]entity.Category, error) {

	query := `
	SELECT id, category_name
	FROM categories
	ORDER BY id
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entity.Category

	for rows.Next() {
		var c entity.Category

		err := rows.Scan(
			&c.ID,
			&c.CategoryName,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, c)
	}

	return categories, nil
}

func (r *categoryRepository) CreateCategory(ctx context.Context, category entity.Category) error {

	query := `
	INSERT INTO categories(category_name)
	VALUES(?)
	`

	_, err := r.db.ExecContext(ctx, query, category.CategoryName)
	if err != nil {
		return errors.New("something went wrong")
	}

	return nil
}

func (r *categoryRepository) AssignCategory(ctx context.Context, pc entity.ProductCategory) error {

	query := `
	INSERT INTO product_categories(product_id, category_id)
	VALUES(?, ?)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		pc.ProductID,
		pc.CategoryID,
	)

	if err != nil {
		return errors.New("something went wrong")
	}

	return nil
}
