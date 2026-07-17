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

type ProductRepo struct {
	ID          int
	SupplierID  int
	ProductName string
	Description string
	Price       float64
	Stock       int
}

type StockReport struct {
	ProductID   int
	ProductName string
	Stock       int
}