package handlers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
)

type adminHandler struct {
	uc domain.AdminUsecase
}

func NewAdminHandler(uc domain.AdminUsecase) domain.AdminHandler {
	return &adminHandler{
		uc: uc,
	}
}

func (h *adminHandler) Run() {
	for {
		fmt.Println("\n===== ADMIN MENU =====")
		fmt.Println("1. List Products")
		fmt.Println("2. Add Product")
		fmt.Println("3. Update Product")
		fmt.Println("4. Delete Product")
		fmt.Println("0. Logout")

		fmt.Print("Choose : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			h.listProducts()
		case 2:
			h.addProduct()
		case 3:
			h.updateProduct()
		case 4:
			h.deleteProduct()
		case 0:
			fmt.Println("Logout...")
			return
		default:
			fmt.Println("Invalid menu")
		}
	}
}

func (h *adminHandler) listProducts() {
	fmt.Println("\n--- List Products ---")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	products, err := h.uc.ListProducts(ctx)
	if err != nil {
		fmt.Printf("Error get products: %v\n", err)
		return
	}

	if len(products) == 0 {
		fmt.Println("No products available.")
		return
	}

	fmt.Printf("%-5s | %-20s | %-15s | %-8s\n", "ID", "Name", "Price", "Stock")
	fmt.Println("-----------------------------------------------------------------")

	for _, p := range products {
		priceFormatted := formatRupiah(p.Price)
		fmt.Printf("%-5d | %-20s | %-15s | %-8d\n", p.ID, p.Name, priceFormatted, p.Stock)
	}
}

func (h *adminHandler) addProduct() {
	fmt.Println("\n--- Add Product ---")
	var name string
	var price float64
	var stock int

	fmt.Print("Product Name  : ")
	fmt.Scanln(&name)
	fmt.Print("Product Price : ")
	fmt.Scanln(&price)
	fmt.Print("Product Stock : ")
	fmt.Scanln(&stock)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	product := entity.Product{
		Name:  name,
		Price: price,
		Stock: stock,
	}

	err := h.uc.AddProduct(ctx, &product)
	if err != nil {
		fmt.Printf("Error add product: %v\n", err)
		return
	}

	fmt.Println("Product added successfully!")
}

func (h *adminHandler) updateProduct() {
	fmt.Println("\n--- Update Product ---")
	var id int
	var name string
	var price float64
	var stock int

	fmt.Print("Product ID to Update : ")
	fmt.Scanln(&id)
	fmt.Print("New Name             : ")
	fmt.Scanln(&name)
	fmt.Print("New Price            : ")
	fmt.Scanln(&price)
	fmt.Print("New Stock            : ")
	fmt.Scanln(&stock)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	product := entity.Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}

	err := h.uc.UpdateProduct(ctx, &product)
	if err != nil {
		fmt.Printf("Error update product: %v\n", err)
		return
	}

	fmt.Println("Product updated successfully!")
}

func (h *adminHandler) deleteProduct() {
	fmt.Println("\n--- Delete Product ---")
	var id int
	fmt.Print("Product ID to Delete : ")
	fmt.Scanln(&id)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.uc.DeleteProduct(ctx, id)
	if err != nil {
		fmt.Printf("Error delete product: %v\n", err)
		return
	}

	fmt.Println("Product deleted successfully!")
}

// Fungsi pembantu untuk format Rupiah
func formatRupiah(amount float64) string {
	s := fmt.Sprintf("%.0f", amount)
	var result []string
	n := len(s)
	for i := n; i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		result = append([]string{s[start:i]}, result...)
	}
	return "Rp " + strings.Join(result, ".")
}
