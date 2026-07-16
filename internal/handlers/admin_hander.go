package handlers

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
)

type adminHandler struct {
	uc     domain.AdminUsecase
	userUC domain.UserUsecase // Ditambahkan untuk akses UpdateProfile
}

// Update constructor agar menerima UserUsecase
func NewAdminHandler(uc domain.AdminUsecase, userUC domain.UserUsecase) domain.AdminHandler {
	return &adminHandler{
		uc:     uc,
		userUC: userUC,
	}
}

func (h *adminHandler) Run() {
	for {
		fmt.Println("\n===== ADMIN MENU =====")
		fmt.Println("1. List Products")
		fmt.Println("2. Add Product")
		fmt.Println("3. Update Product")
		fmt.Println("4. Delete Product")
<<<<<<< HEAD
		fmt.Println("5. List Orders")
		fmt.Println("6. Update Status Orders to Completed")
		fmt.Println("7. Report Completed Orders")
		fmt.Println("8. Assign Category to Product")
		fmt.Println("9. User Report - Paling Banyak Belanja")
		fmt.Println("10. Stock Report - Stok Habis/0")
		fmt.Println("11. Update Profile") // Menu baru
=======
		fmt.Println("\n======================")
		fmt.Println("5. List Orders")
		fmt.Println("6. Update Status Orders to Completed")
		fmt.Println("\n======================")
		fmt.Println("7. Report Completed Orders")
>>>>>>> a9d2308fdc4245458fd69dd2b7b286b0217a42fc
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
		case 5:
<<<<<<< HEAD
			fmt.Println("\n[Menu ini sedang dikerjakan oleh Kesaa]")
		case 6:
			fmt.Println("\n[Menu ini sedang dikerjakan oleh Kesaa]")
		case 7:
			fmt.Println("\n[Menu ini sedang dikerjakan oleh Kesaa]")
		case 8:
			h.assignCategoryMenu()
		case 9:
			h.showUserReport()
		case 10:
			h.showStockReport()
		case 11:
			h.updateProfile() // Memanggil fungsi baru
=======
			h.listOrders()
		case 6:
			h.updateOrderStatusToCompleted()
		case 7:
			h.reportCompletedOrder()
>>>>>>> a9d2308fdc4245458fd69dd2b7b286b0217a42fc
		case 0:
			fmt.Println("Logout...")
			return
		default:
			fmt.Println("Invalid menu")
		}
	}
}

// --- Fungsi Baru: Update Profile ---
func (h *adminHandler) updateProfile() {
	fmt.Println("\n--- Update Profile ---")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter User ID to Update : ")
	var userID int
	fmt.Scanln(&userID)

	fmt.Print("Enter Full Name         : ")
	fullName, _ := reader.ReadString('\n')
	fullName = strings.TrimSpace(fullName)

	fmt.Print("Enter Company Name      : ")
	companyName, _ := reader.ReadString('\n')
	companyName = strings.TrimSpace(companyName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.userUC.UpdateProfile(ctx, entity.UserProfile{
		UserID:      userID,
		FullName:    fullName,
		CompanyName: companyName,
	})

	if err != nil {
		fmt.Printf("Error updating profile: %v\n", err)
		return
	}

	fmt.Println("Profile updated successfully!")
}

// --- Fungsi existing lainnya tetap di bawah sini ---

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

<<<<<<< HEAD
func (h *adminHandler) assignCategoryMenu() {
	fmt.Println("\n--- Assign Category to Product ---")
	var productID int
	var categoryID int

	fmt.Print("Product ID   : ")
	fmt.Scanln(&productID)
	fmt.Print("Category ID  : ")
	fmt.Scanln(&categoryID)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.uc.AssignCategoryToProduct(ctx, productID, categoryID)
	if err != nil {
		fmt.Printf("Error assigning category: %v\n", err)
		return
	}

	fmt.Println("Category assigned successfully!")
}

func (h *adminHandler) showUserReport() {
	fmt.Println("\n--- User Report (Most Active Buyers) ---")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reports, err := h.uc.GetUserReport(ctx)
	if err != nil {
		fmt.Printf("Error get user report: %v\n", err)
		return
	}

	if len(reports) == 0 {
		fmt.Println("No user reports available.")
		return
	}

	fmt.Printf("%-5s | %-25s | %-20s | %-15s | %-12s\n", "ID", "Email Customer", "Nama Perusahaan", "Nama Kontak", "Total Order")
	fmt.Println("---------------------------------------------------------------------------------------------------")

	for _, ur := range reports {
		fmt.Printf("%-5d | %-25s | %-20s | %-15s | %-12d\n", ur.ID, ur.Email, ur.CompanyName, ur.ContactName, ur.TotalOrders)
	}
}

func (h *adminHandler) showStockReport() {
	fmt.Println("\n--- Stock Report (Out of Stock / 0) ---")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reports, err := h.uc.GetStockReport(ctx)
	if err != nil {
		fmt.Printf("Error get stock report: %v\n", err)
		return
	}

	if len(reports) == 0 {
		fmt.Println("No products out of stock (Stock = 0).")
		return
	}

	fmt.Printf("%-5s | %-25s | %-10s | %-15s\n", "ID", "Product Name", "Stock", "Price")
	fmt.Println("-----------------------------------------------------------------------------")

	for _, sr := range reports {
		priceFormatted := formatRupiah(sr.Price)
		fmt.Printf("%-5d | %-25s | %-10d | %-15s\n", sr.ID, sr.ProductName, sr.Stock, priceFormatted)
=======
func (h *adminHandler) listOrders() {

	ctx := context.Background()

	orders, err := h.uc.ListOrders(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(orders) == 0 {
		fmt.Println("No orders found.")
		return
	}

	fmt.Println("\n========================================== ALL ORDERS ==========================================")
	fmt.Printf("%-5s %-30s %-12s %-12s %-20s\n",
		"ID",
		"Email",
		"Total",
		"Status",
		"Order Date",
	)

	fmt.Println("-----------------------------------------------------------------------------------------------")

	for _, order := range orders {

		fmt.Printf("%-5d %-30s Rp%-10.0f %-12s %-20s\n",
			order.ID,
			order.Email,
			order.TotalPrice,
			order.Status,
			order.CreatedAt.Format("2006-01-02 15:04:05"),
		)
	}

	fmt.Println("-----------------------------------------------------------------------------------------------")
}

func (h *adminHandler) updateOrderStatusToCompleted() {

	ctx := context.Background()

	orders, err := h.uc.GetOrdersByStatus(ctx, entity.StatusPending)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(orders) == 0 {
		fmt.Println("No pending orders.")
		return
	}

	fmt.Println("\n==================== PENDING ORDERS ====================")
	fmt.Printf("%-8s %-30s %-12s %-12s\n",
		"ID",
		"Email",
		"Total",
		"Status",
	)

	for _, order := range orders {

		fmt.Printf("%-8d %-30s Rp%-10.0f %-12s\n",
			order.ID,
			order.Email,
			order.TotalPrice,
			order.Status,
		)
	}

	var orderID int

	fmt.Print("\nEnter Order ID to Complete : ")
	fmt.Scanln(&orderID)

	// Change to completed
	err = h.uc.UpdateOrderStatus(ctx, orderID, entity.StatusCompleted)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Order updated successfully.")
}

func (h *adminHandler) reportCompletedOrder() {

	ctx := context.Background()

	orders, err := h.uc.GetOrdersByStatus(ctx, entity.StatusCompleted)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(orders) == 0 {
		fmt.Println("No pending orders.")
		return
	}

	fmt.Println("\n==================== COMPLETED ORDERS ====================")
	fmt.Printf("%-8s %-30s %-12s %-12s\n",
		"ID",
		"Email",
		"Total",
		"Status",
	)

	for _, order := range orders {

		fmt.Printf("%-8d %-30s Rp%-10.0f %-12s\n",
			order.ID,
			order.Email,
			order.TotalPrice,
			order.Status,
		)
>>>>>>> a9d2308fdc4245458fd69dd2b7b286b0217a42fc
	}
}
