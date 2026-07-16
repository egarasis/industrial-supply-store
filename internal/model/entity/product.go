package entity

import "time"

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductWithSupplier struct {
	ID           int
	SupplierID   int
	SupplierName string
	ProductName  string
	Description  string
	Price        float64
	Stock        int
}

type CartItem struct {
	ProductID   int
	ProductName string
	Price       float64
	Quantity    int
	Subtotal    float64
}

type UserReport struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	CompanyName string `json:"company_name"`
	ContactName string `json:"contact_name"`
	TotalOrders int    `json:"total_orders"`
}

type StockReport struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}
