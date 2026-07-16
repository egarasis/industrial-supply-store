package db

import (
	"context"
	"database/sql"
	"errors"

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
	// Menggunakan 'name' sesuai struktur database asli
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
	// Menggunakan 'name' sesuai struktur database asli
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
	// Menggunakan 'name' sesuai struktur database asli
	query := "UPDATE products SET name = ?, price = ?, stock = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, product.Name, product.Price, product.Stock, product.ID)
	return err
}

func (r *productRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *productRepository) GetAllProducts(ctx context.Context) ([]entity.ProductWithSupplier, error) {
	// Di-alias 'name AS product_name' agar cocok dengan struct ProductWithSupplier milik tim
	query := `
    SELECT 
        p.id, 
        p.supplier_id, 
        s.supplier_name, 
        p.name AS product_name, 
        p.description, 
        p.price, 
        p.stock
    FROM products p
    JOIN suppliers s 
    ON p.supplier_id = s.id
    ORDER BY p.id;
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.ProductWithSupplier
	for rows.Next() {
		var product entity.ProductWithSupplier
		err := rows.Scan(
			&product.ID,
			&product.SupplierID,
			&product.SupplierName,
			&product.ProductName,
			&product.Description,
			&product.Price,
			&product.Stock,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *productRepository) GetProductByID(ctx context.Context, id int) (entity.ProductWithSupplier, error) {
	// Di-alias 'name AS product_name' agar cocok dengan struct ProductWithSupplier milik tim
	query := `
    SELECT 
        p.id, 
        p.supplier_id, 
        s.supplier_name, 
        p.name AS product_name, 
        p.description, 
        p.price, 
        p.stock
    FROM products p
    JOIN suppliers s 
    ON p.supplier_id = s.id
    WHERE p.id = ?;
    `

	var product entity.ProductWithSupplier
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.SupplierID,
		&product.SupplierName,
		&product.ProductName,
		&product.Description,
		&product.Price,
		&product.Stock,
	)

	if err != nil {
		err = errors.New("something went wrong. Please try again later.")
	}

	return product, err
}

// =====================
// ADMIN
// =====================

func (r *productRepository) CreateProduct(ctx context.Context, product entity.ProductWithSupplier) error {
	// Disesuaikan ke kolom 'name'
	query := `
    INSERT INTO products 
    (
        supplier_id, 
        name, 
        description, 
        price, 
        stock
    ) 
    VALUES 
    (
        ?, ?, ?, ?, ?
    )
    `

	_, err := r.db.ExecContext(
		ctx,
		query,
		product.SupplierID,
		product.ProductName, // di-map dari properti struct
		product.Description,
		product.Price,
		product.Stock,
	)

	if err != nil {
		return errors.New("something went wrong. Please try again later.")
	}

	return nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, product entity.ProductWithSupplier) error {
	// Disesuaikan ke kolom 'name'
	query := `
    UPDATE products 
    SET 
        supplier_id = ?, 
        name = ?, 
        description = ?, 
        price = ?, 
        stock = ? 
    WHERE id = ?;
    `

	_, err := r.db.ExecContext(
		ctx,
		query,
		product.SupplierID,
		product.ProductName,
		product.Description,
		product.Price,
		product.Stock,
		product.ID,
	)

	if err != nil {
		return errors.New("something went wrong. Please try again later.")
	}

	return nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, id int) error {
	query := `
    DELETE FROM products 
    WHERE id = ?;
    `

	_, err := r.db.ExecContext(ctx, query, id)
<<<<<<< HEAD
	return err
=======
	if err != nil {
		return errors.New("something went wrong. Please try again later.")
	}

	return nil
>>>>>>> a9d2308fdc4245458fd69dd2b7b286b0217a42fc
}

// =====================
// CHECKOUT
// =====================

func (r *productRepository) UpdateStock(ctx context.Context, tx *sql.Tx, productID int, qty int) error {
	query := `
    UPDATE products 
    SET stock = stock - ? 
    WHERE id = ? 
    AND stock >= ?;
    `

	result, err := tx.ExecContext(
		ctx,
		query,
		qty,
		productID,
		qty,
	)
	if err != nil {
		return errors.New("something went wrong. Please try again later.")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("something went wrong. Please try again later")
	}

	if rowsAffected == 0 {
		return errors.New("stock is not enough")
	}

	return nil
}
<<<<<<< HEAD

// =========================================================================
// JATAH TUGAS SAYA: QUERY ASSIGN CATEGORY, USER REPORT, & STOCK REPORT
// =========================================================================

func (r *productRepository) AssignCategory(ctx context.Context, productID, categoryID int) error {
	query := `INSERT INTO product_categories (product_id, category_id) VALUES (?, ?)`
	_, err := r.db.ExecContext(ctx, query, productID, categoryID)
	return err
}

func (r *productRepository) GetUserReport(ctx context.Context) ([]entity.UserReport, error) {
	query := `SELECT u.id, u.email, IFNULL(up.company_name, ''), IFNULL(up.contact_name, ''), COUNT(o.id) AS total_orders
              FROM users u
              LEFT JOIN user_profiles up ON u.id = up.user_id
              LEFT JOIN orders o ON u.id = o.user_id
              WHERE u.role = 'customer'
              GROUP BY u.id, up.company_name, up.contact_name
              ORDER BY total_orders DESC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []entity.UserReport
	for rows.Next() {
		var ur entity.UserReport
		if err := rows.Scan(&ur.ID, &ur.Email, &ur.CompanyName, &ur.ContactName, &ur.TotalOrders); err != nil {
			return nil, err
		}
		reports = append(reports, ur)
	}
	return reports, nil
}

func (r *productRepository) GetStockReport(ctx context.Context) ([]entity.StockReport, error) {
	// Di-alias 'name AS product_name' agar aman dibaca database lama
	query := `SELECT id, name AS product_name, stock, price FROM products WHERE stock = 0`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []entity.StockReport
	for rows.Next() {
		var sr entity.StockReport
		if err := rows.Scan(&sr.ID, &sr.ProductName, &sr.Stock, &sr.Price); err != nil {
			return nil, err
		}
		reports = append(reports, sr)
	}
	return reports, nil
}
=======
>>>>>>> a9d2308fdc4245458fd69dd2b7b286b0217a42fc
