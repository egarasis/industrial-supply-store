package db

import (
	"context"
	"database/sql"

	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
)

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) domain.OrderRepository {
	return &orderRepository{
		db: db,
	}
}

// =====================
// CHECKOUT
// =====================

func (r *orderRepository) CreateOrder(
	ctx context.Context,
	tx *sql.Tx,
	order entity.Order,
) (int, error) {

	query := `
	INSERT INTO orders
	(
		user_id,
		total_price,
		status
	)
	VALUES
	(
		?,
		?,
		?
	)
	`

	result, err := tx.ExecContext(
		ctx,
		query,
		order.UserID,
		order.TotalPrice,
		order.Status,
	)

	if err != nil {
		return 0, err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(orderID), nil
}

func (r *orderRepository) CreateOrderItem(
	ctx context.Context,
	tx *sql.Tx,
	item entity.OrderItem,
) error {

	query := `
	INSERT INTO order_items
	(
		order_id,
		product_id,
		quantity,
		subtotal
	)
	VALUES
	(
		?,
		?,
		?,
		?
	);
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		item.OrderID,
		item.ProductID,
		item.Quantity,
		item.Subtotal,
	)

	return err
}

func (r *orderRepository) UpdateOrderTotal(
	ctx context.Context,
	tx *sql.Tx,
	orderID int,
	total float64,
) error {

	query := `
	UPDATE orders
	SET total_price = ?
	WHERE id = ?;
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		total,
		orderID,
	)

	return err
}

// =====================
// CUSTOMER
// =====================

func (r *orderRepository) GetOrdersByUserID(
	ctx context.Context,
	userID int,
) ([]entity.Order, error) {

	query := `
	SELECT
		id,
		user_id,
		total_price,
		status,
		created_at
	FROM orders
	WHERE user_id = ?
	ORDER BY created_at DESC;
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order

	for rows.Next() {

		var order entity.Order

		err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.TotalPrice,
			&order.Status,
			&order.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *orderRepository) GetOrderByID(
	ctx context.Context,
	orderID int,
) (entity.Order, error) {

	query := `
	SELECT
		id,
		user_id,
		total_price,
		status,
		created_at
	FROM orders
	WHERE id = ?;
	`

	var order entity.Order

	err := r.db.QueryRowContext(
		ctx,
		query,
		orderID,
	).Scan(
		&order.ID,
		&order.UserID,
		&order.TotalPrice,
		&order.Status,
		&order.CreatedAt,
	)

	return order, err
}

func (r *orderRepository) GetOrderItems(
	ctx context.Context,
	orderID int,
) ([]entity.OrderItem, error) {

	query := `
	SELECT
		p.product_name,
		p.price,
		oi.quantity,
		oi.subtotal
	FROM order_items oi
	JOIN products p
	ON oi.product_id = p.id
	WHERE oi.order_id = ?;
	`

	rows, err := r.db.QueryContext(ctx, query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var details []entity.OrderItem

	for rows.Next() {

		var detail entity.OrderItem

		err := rows.Scan(
			&detail.ProductName,
			&detail.Price,
			&detail.Quantity,
			&detail.Subtotal,
		)

		if err != nil {
			return nil, err
		}

		details = append(details, detail)
	}

	return details, nil
}

// =====================
// ADMIN
// =====================

func (r *orderRepository) GetAllOrders(
	ctx context.Context,
) ([]entity.Order, error) {

	query := `
	SELECT
		id,
		user_id,
		total_price,
		status,
		created_at
	FROM orders
	ORDER BY created_at DESC;
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order

	for rows.Next() {

		var order entity.Order

		err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.TotalPrice,
			&order.Status,
			&order.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *orderRepository) UpdateOrderStatus(
	ctx context.Context,
	orderID int,
	status string,
) error {

	query := `
	UPDATE orders
	SET status = ?
	WHERE id = ?;
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		status,
		orderID,
	)

	return err
}
