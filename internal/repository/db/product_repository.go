package db

import (
	"context"
	"database/sql"

	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) FindAll(ctx context.Context) ([]entity.Product, error) {
	query := "SELECT id, name, price, stock, created_at FROM products"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var p entity.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *productRepository) Create(ctx context.Context, product *entity.Product) error {
	query := "INSERT INTO products (name, price, stock) VALUES (?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, product.Name, product.Price, product.Stock)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	product.ID = int(id)
	return nil
}

func (r *productRepository) Update(ctx context.Context, product *entity.Product) error {
	query := "UPDATE products SET name = ?, price = ?, stock = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, product.Name, product.Price, product.Stock, product.ID)
	return err
}

func (r *productRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
