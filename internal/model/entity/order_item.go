package entity

type OrderItem struct {
	ID          int
	OrderID     int
	ProductID   int
	ProductName string
	Quantity    int
	Price       float64
	Subtotal    float64
}