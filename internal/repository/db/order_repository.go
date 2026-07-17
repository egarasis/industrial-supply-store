package db

import (
	"context"
	"database/sql"
	"errors"

	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
)

type orderRepository struct {
	db          *sql.DB
	productRepo domain.ProductRepository
}

func NewOrderRepository(
	db *sql.DB,
	productRepo domain.ProductRepository,
) domain.OrderRepository {

	return &orderRepository{
		db:          db,
		productRepo: productRepo,
	}
}

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
	orderID, userID int,
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
	JOIN orders o
		ON oi.order_id = o.id
	WHERE oi.order_id = ?
		AND o.user_id = ?;
	`

	rows, err := r.db.QueryContext(ctx, query, orderID, userID)
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

func (r *orderRepository) GetAllOrders(
	ctx context.Context,
) ([]entity.OrderWithUser, error) {

	query := `
	SELECT
		o.id,
		o.user_id,
		u.email,
		o.total_price,
		o.status,
		o.created_at
	FROM orders o
	JOIN users u
		ON o.user_id = u.id
	ORDER BY o.created_at DESC;
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.OrderWithUser

	for rows.Next() {

		var order entity.OrderWithUser

		err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.Email,
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

func (r *orderRepository) GetOrdersByStatus(
	ctx context.Context,
	status string,
) ([]entity.OrderWithUser, error) {

	query := `
	SELECT
		o.id,
		o.user_id,
		u.email,
		o.total_price,
		o.status,
		o.created_at
	FROM orders o
	JOIN users u
		ON o.user_id = u.id
	WHERE o.status = ?
	ORDER BY o.created_at ASC;
	`

	rows, err := r.db.QueryContext(ctx, query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.OrderWithUser

	for rows.Next() {

		var order entity.OrderWithUser

		err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.Email,
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

func (r *orderRepository) Checkout(
	ctx context.Context,
	userID int,
	cart []entity.CartItem,
) error {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	order := entity.Order{
		UserID:     userID,
		TotalPrice: 0,
		Status:     "Pending",
	}

	orderID, err := r.CreateOrder(ctx, tx, order)
	if err != nil {
		return err
	}

	var total float64

	for _, cartItem := range cart {

		product, err := r.productRepo.GetProductByID(ctx, cartItem.ProductID)
		if err != nil {
			return err
		}

		if product.Stock < cartItem.Quantity {
			return errors.New("stock is not enough")
		}

		subtotal := product.Price * float64(cartItem.Quantity)

		item := entity.OrderItem{
			OrderID:   orderID,
			ProductID: product.ID,
			Quantity:  cartItem.Quantity,
			Subtotal:  subtotal,
		}

		if err := r.CreateOrderItem(ctx, tx, item); err != nil {
			return err
		}

		if err := r.productRepo.UpdateStock(ctx, tx, product.ID, cartItem.Quantity); err != nil {
			return err
		}

		total += subtotal
	}

	if err := r.UpdateOrderTotal(ctx, tx, orderID, total); err != nil {
		return err
	}

	return tx.Commit()
}
