package entity

import "time"

type Order struct {
	ID         int
	UserID     int
	TotalPrice float64
	Status     string
	CreatedAt  time.Time
}

// For join
type OrderDetail struct {
	ProductName string
	Price       float64
	Quantity    int
	Subtotal    float64
}