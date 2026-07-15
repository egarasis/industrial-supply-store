package handlers

import (
	"fmt"

	"industrial-supply-store/internal/domain"
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
	fmt.Println("List Product")
}

func (h *adminHandler) addProduct() {
	fmt.Println("Add Product")
}

func (h *adminHandler) updateProduct() {
	fmt.Println("Update Product")
}

func (h *adminHandler) deleteProduct() {
	fmt.Println("Delete Product")
}