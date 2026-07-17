package entity

import "time"

const (
	StatusPending   = "Pending"
	StatusCompleted = "Completed"
)

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

// Orders join with users
type OrderWithUser struct {
	ID         int
	UserID     int
	Email      string
	TotalPrice float64
	Status     string
	CreatedAt  time.Time
}

type UserReport struct {
	UserID     int
	Email      string
	TotalOrder int
	TotalSpent float64
}
