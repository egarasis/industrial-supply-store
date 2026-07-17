package entity

type Category struct {
	ID           int
	CategoryName string
}

type ProductCategory struct {
	ProductID  int
	CategoryID int
}
